package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	hexKey := "your-aes-key"
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		panic(err)
	}

	transactionId := []byte("your-trx-id")

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, transactionId, nil)

	tagSize := aesgcm.Overhead()
	tag := ciphertext[len(ciphertext)-tagSize:]
	ciphertextWithoutTag := ciphertext[:len(ciphertext)-tagSize]

	fmt.Println("Ciphertext (hex):", hex.EncodeToString(ciphertextWithoutTag))
	fmt.Println("Nonce (hex):     ", hex.EncodeToString(nonce))
	fmt.Println("Tag (hex):       ", hex.EncodeToString(tag))

}
