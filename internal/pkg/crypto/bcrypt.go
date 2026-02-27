package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func ValidatePassword(hashed string, passwowrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwowrd))
	return err == nil
}
