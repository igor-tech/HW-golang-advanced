package user

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func GenerateSessionID() (string, error) {
	const bytesLen = 16
	bytes := make([]byte, bytesLen)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func GenerateCode() (string, error) {
	const max = 10000 // 0000…9999
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%04d", nBig.Int64()), nil
}

func HashCode(code, salt string) string {
	h := sha256.Sum256([]byte(code + salt))
	return hex.EncodeToString(h[:])
}

func CheckHash(code, sessionID, hash string) bool {
	return HashCode(code, sessionID) == hash
}
