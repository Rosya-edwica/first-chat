package util

import "errors"

func HashPassword(password string) (string, error) {
	// Невозможно скачать библиотеку из-за VPN "golang.org/x/crypto/bcrypt"

	return password + "hash", nil
}

func CheckPassword(password string, hashedPassword string) error {
	if password+"hash" == hashedPassword {
		return nil
	}
	return errors.New("Неправильный пароль!")
}
