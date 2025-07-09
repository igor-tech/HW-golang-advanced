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

func GenerateCode() (int, error) {
	const max = 10000 // 0000…9999
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}

	return int(nBig.Int64()), nil
}

func HashCode(code int, salt string) string {
	h := sha256.Sum256([]byte(fmt.Sprintf("%d%s", code, salt)))
	return hex.EncodeToString(h[:])
}

func CheckHash(code int, sessionID, hash string) bool {
	return HashCode(code, sessionID) == hash
}
