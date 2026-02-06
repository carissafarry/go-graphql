package user

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repo            Repository
	pendingUserRepo PendingUserStore
	otpRepo         OTPStore
	otpGen          OTPGenerator
	validator       Validator
}

func NewUsecase(
	repo Repository,
	pendingUserRepo PendingUserStore,
	otpRepo OTPStore,
	otpGen OTPGenerator,
	validator Validator,
) *Usecase {
	return &Usecase{
		repo:            repo,
		pendingUserRepo: pendingUserRepo,
		otpRepo:         otpRepo,
		otpGen:          otpGen,
		validator:       validator,
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

	if err := u.validator.ValidateCreateUser(fullName, email); err != nil {
		return nil, err
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

	if err := u.validator.ValidateRegister(email, password); err != nil {
		return err
	}

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

	if err := u.validator.ValidateVerifyOTP(email, otp); err != nil {
		return err
	}

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