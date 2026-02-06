package user

type validator struct{}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) ValidateCreateUser(fullName, email string) error {
	if fullName == "" {
		return ErrFullNameRequired
	}
	if err := v.requireEmail(email); err != nil {
		return err
	}
	return nil
}

func (v *validator) ValidateRegister(email, password string) error {
	if err := v.requireEmail(email); err != nil {
		return err
	}
	if err := v.requirePassword(password); err != nil {
		return err
	}
	if len(password) < 8 {
		return ErrInvalidPassword
	}
	return nil
}

func (v *validator) ValidateVerifyOTP(email, otp string) error {
	if err := v.requireEmail(email); err != nil {
		return err
	}
	if otp == "" {
		return ErrInvalidOTP
	}
	return nil
}


// ===== PRIVATE HELPERS =====

func (v *validator) requireEmail(email string) error {
	if email == "" {
		return ErrEmailRequired
	}
	return nil
}

func (v *validator) requirePassword(password string) error {
	if password == "" {
		return ErrInvalidPassword
	}
	return nil
}