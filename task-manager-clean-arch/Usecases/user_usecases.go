package usecases

import (
	"errors"

	"task-manager-clean-arch/Domain"
)

type UserUseCase struct {
	userRepo        domain.UserRepository
	passwordService domain.PasswordService
	jwtService      domain.JWTService
}

func NewUserUseCase(userRepo domain.UserRepository, passwordService domain.PasswordService, jwtService domain.JWTService) *UserUseCase {
	return &UserUseCase{
		userRepo:        userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (uc *UserUseCase) RegisterUser(username, password string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters long")
	}

	exists, err := uc.userRepo.ExistsByUsername(username)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("username already taken")
	}

	hashedPassword, err := uc.passwordService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	userCount, err := uc.userRepo.GetUserCount()
	if err != nil {
		return nil, err
	}

	role := domain.RoleUser
	if userCount == 0 {
		role = domain.RoleAdmin
	}

	user := &domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) LoginUser(username, password string) (string, error) {
	if username == "" {
		return "", errors.New("username is required")
	}

	if password == "" {
		return "", errors.New("password is required")
	}

	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !uc.passwordService.ComparePassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := uc.jwtService.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UserUseCase) ValidateToken(tokenString string) (*domain.JWTClaims, error) {
	if tokenString == "" {
		return nil, errors.New("token is required")
	}

	claims, err := uc.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
} 