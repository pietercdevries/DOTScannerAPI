package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func Encrypt(passPhrase string) string {
	text := []byte(passPhrase)
	key := []byte(getSecretCodeFromFile())

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, text, nil))
}

func Decrypt(encodedString string) string{
	key := []byte(getSecretCodeFromFile())
	var ciphertext, _ = hex.DecodeString(encodedString)

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)
}

func getSecretCodeFromFile() string {
	content, err := ioutil.ReadFile("cred.key")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	return text
}
