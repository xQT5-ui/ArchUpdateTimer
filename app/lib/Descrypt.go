package lib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	log "app.go/app/lib/logger"
)

func encrypt(password string) string {
	// Calculate SHA256 hash
	hash := sha256.Sum256([]byte(password))

	// Convert hash to string
	return hex.EncodeToString(hash[:])
}

func decrypt(hash string, lg *log.Logger) string {
	// Convert base64 string to bytes
	bytes, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		lg.Error(err, "Ошибка декодирования пароля:")
		return ""
	}

	return string(bytes)
}

func GetDecryptPassword(password, hash string, lg *log.Logger) string {
	hash_config := encrypt(password)

	if hash_config == hash {
		lg.Info("Пароль совпадает")
		return decrypt(password, lg)
	} else {
		lg.Error(fmt.Errorf("неверное совпадение пароля"), "Ошибка проверки пароля")
		return ""
	}
}
