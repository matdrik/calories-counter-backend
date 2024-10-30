package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckParameters todo - сделать нормальную проверку параметров
func CheckParameters(username, usernameDB, password, passwordHash string) error {
	if username != usernameDB || !CheckPasswordHash(password, passwordHash) {
		return errors.New("неверные данные пользователя")
	}
	return nil
}
