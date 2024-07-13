package lib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"log"
)

func encrypt(password string) string {
	// Calculate SHA256 hash
	hash := sha256.Sum256([]byte(password))

	// Convert hash to string
	return hex.EncodeToString(hash[:])
}

func decrypt(hash string) string {
	// Convert base64 string to bytes
	bytes, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		log.Fatalf("ошибка декодирования пароля: %v", err)
		return ""
	}
	return string(bytes)
}

func GetDecryptPassword(password, hash string) string {
	hash_config := encrypt(password)

	if hash_config == hash {
		return decrypt(password)
	} else {
		log.Fatal("неверное совпадение пароля")
		return ""
	}
}
