package salter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"os"

	"github.com/senonerk/sup/shared/aerr"
)

var hmc hash.Hash

// GenerateHMAC generates HMAC string
func GenerateHMAC(password string) (string, error) {
	if err := loadSecret(); err != nil {
		return "", err
	}

	hmc.Reset()

	if _, err := hmc.Write([]byte(password)); err != nil {
		return "", err
	}

	return hex.EncodeToString(hmc.Sum(nil)), nil
}

// CompareHMAC compares password with its HMAC hash
func CompareHMAC(password, hmacString string) error {
	if err := loadSecret(); err != nil {
		return err
	}

	hash, err := GenerateHMAC(password)
	if err != nil {
		return err
	}

	if !hmac.Equal([]byte(hash), []byte(hmacString)) {
		return &aerr.AppError{
			Code:    400,
			Message: "Incorrect password",
			Source:  "senonerk.sup.srv.auth.salter",
		}
	}

	return nil
}

func loadSecret() error {
	if hmc != nil {
		return nil
	}

	if secret := os.Getenv("PASSWORD_HMAC_SECRET"); secret != "" {
		hmc = hmac.New(sha256.New, []byte(secret))
		return nil
	}

	return errors.New("Password HMAC secret not set")
}
