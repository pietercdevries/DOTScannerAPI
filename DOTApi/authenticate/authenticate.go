package authenticate

import (
	"crypto/rand"
	"math/big"
)

func GenerateRefreshToken() string {
	return generateNewToken(16)
}

func GenerateAuthenticateToken() string {
	return generateNewToken(26)
}

func generateNewToken(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

