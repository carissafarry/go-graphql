package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repo            Repository
	pendingUserRepo PendingUserStore
	otpRepo         OTPStore
	otpGen          OTPGenerator
}

type OTPGenerator interface {
	GenerateOTP() (string, error)
}

func NewUsecase(
	repo Repository,
	pendingUserRepo PendingUserStore,
	otpRepo OTPStore,
	otpGen OTPGenerator,
) *Usecase {
	return &Usecase{
		repo:            repo,
		pendingUserRepo: pendingUserRepo,
		otpRepo:         otpRepo,
		otpGen:          otpGen,
	}
}

func (u *Usecase) GetUsers(ctx context.Context) ([]*User, error) {
	return u.repo.FindAll(ctx)
}

func (u *Usecase) CreateUser(
	ctx context.Context,
	fullName string,
	email string,
) (*User, error) {
	// Validation
	if strings.TrimSpace(fullName) == "" {
		return nil, errors.New("fullName is required")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}

	user := &User{
		FullName: fullName,
		Email:    email,
	}

	// Orchestration
	createdUser, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *Usecase) Register(
	ctx context.Context,
	email string,
	password string,
) error {

	existing, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existing != nil {
		return ErrEmailAlreadyRegistered
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pendingUser := &PendingUser{
		Email:    email,
		Password: string(hashed),
	}

	if err := u.pendingUserRepo.Save(
		ctx,
		pendingUser,
		30*time.Minute,
	); err != nil {
		return err
	}

	otp, err := u.otpGen.GenerateOTP()
	if err != nil {
		return err
	}

	if err := u.otpRepo.Save(
		ctx,
		email,
		otp,
		5*time.Minute,
	); err != nil {
		return err
	}

	return nil
}

func (u *Usecase) VerifyOTP(
	ctx context.Context,
	email string,
	otp string,
) error {

	savedOTP, err := u.otpRepo.Find(ctx, email)
	if err != nil || savedOTP != otp {
		return ErrInvalidOTP
	}

	pending, err := u.pendingUserRepo.Find(ctx, email)
	if err != nil || pending == nil {
		return ErrInvalidOTP
	}

	user := &User{
		Email:    pending.Email,
		Password: pending.Password,
	}

	_, err = u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	_ = u.otpRepo.Delete(ctx, email)
	_ = u.pendingUserRepo.Delete(ctx, email)

	return nil
}