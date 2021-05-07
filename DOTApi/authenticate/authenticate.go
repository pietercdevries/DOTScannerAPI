package authenticate

import (
	"crypto/rand"
	"math/big"
)

func RefreshToken(userId int64) string {
	return generateNewToken(userId)
}

func AuthenticateToken(token string) string {
	return ""
}

func generateNewToken(userId int64) string {
	return generateRandomString(16)
}

func generateRandomString(n int) string {
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

