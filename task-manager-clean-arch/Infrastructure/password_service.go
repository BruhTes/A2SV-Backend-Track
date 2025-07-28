package infrastructure

import (
	"task-manager-clean-arch/Domain"

	"golang.org/x/crypto/bcrypt"
)

type PasswordServiceImpl struct{}

func NewPasswordService() domain.PasswordService {
	return &PasswordServiceImpl{}
}

func (p *PasswordServiceImpl) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (p *PasswordServiceImpl) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
} 