package utils

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func SecureHash(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hashed), nil
}

func CheckSecureHash(noHashed string, hashed string) error {
	hashedPass, err := hex.DecodeString(hashed)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(hashedPass, []byte(noHashed))
}
