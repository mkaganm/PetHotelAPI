package auth

import (
	"golang.org/x/crypto/bcrypt"
)

const StaticHash = "$2a$14$3IBxFqVM3JmRobQHLqVLY.prUhvJgtwEshS007n6hYoTxgD1lLRxe"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckAPIKey(password string) bool {
	hash := StaticHash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckHashPassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
