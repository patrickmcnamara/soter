package crypto

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh/terminal"
)

func GenerateKeypair() (ed25519.PublicKey, ed25519.PrivateKey) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	return publicKey, privateKey
}

func Sign(data []byte, privateKey ed25519.PrivateKey) (signature []byte) {
	signature = ed25519.Sign(privateKey, data)
	return signature
}

func Verify(data, signature []byte, publicKey ed25519.PublicKey) (verification bool) {
	verification = ed25519.Verify(publicKey, data, signature)
	return verification
}

func GetPublicKey(privateKey ed25519.PrivateKey) (publicKey ed25519.PublicKey) {
	publicKey = []byte(privateKey[32:])
	return publicKey
}

func EncryptKey(privateKey ed25519.PrivateKey) (ciphertext []byte) {
	password := getPassword()
	encryptionKey := getEncryptionKey(password)
	aead, _ := chacha20poly1305.New(encryptionKey)
	ciphertext = aead.Seal(ciphertext, make([]byte, 12), privateKey, nil)
	return ciphertext
}

func DecryptKey(ciphertext []byte) (privateKey ed25519.PrivateKey) {
	password := getPassword()
	encryptionKey := getEncryptionKey(password)
	aead, _ := chacha20poly1305.New(encryptionKey)
	privateKey, _ = aead.Open(privateKey, make([]byte, 12), ciphertext, nil)
	return privateKey
}

func getEncryptionKey(password []byte) (encryptionKey []byte) {
	encryptionKey = argon2.IDKey(password, nil, 1, 64*1024, 4, 32)
	return encryptionKey
}

func getPassword() (password []byte) {
	fmt.Print("Please enter your password: ")
	password, _ = terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n\n")
	return password
}
