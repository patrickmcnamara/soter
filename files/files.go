package files

import (
	"io/ioutil"

	"github.com/patrickmcnamara/soter/crypto"

	"golang.org/x/crypto/ed25519"
)

func WriteKeyfile(keyfile string, privateKey ed25519.PrivateKey) {
	privateKeyEncrypted := crypto.EncryptKey(privateKey)
	ioutil.WriteFile(keyfile, privateKeyEncrypted, 0600)
}

func ReadKeyfile(keyfile string) (privateKey ed25519.PrivateKey) {
	privateKeyEncrypted, _ := ioutil.ReadFile(keyfile)
	privateKey = crypto.DecryptKey(privateKeyEncrypted)
	return privateKey
}

func ChangeKeyfilePassword(keyfile string) {
	privateKeyEncrypted, _ := ioutil.ReadFile(keyfile)
	privateKey := crypto.DecryptKey(privateKeyEncrypted)
	privateKeyEncrypted = crypto.EncryptKey(privateKey)
	ioutil.WriteFile(keyfile, privateKeyEncrypted, 0600)
}

func BackupKeyfile(file, keyfile string) {
	data, _ := ioutil.ReadFile(keyfile)
	ioutil.WriteFile(file, data, 0600)
}

func RestoreKeyfile(file, keyfile string) {
	data, _ := ioutil.ReadFile(file)
	ioutil.WriteFile(keyfile, data, 0600)
}

func SignFile(privateKey ed25519.PrivateKey, file string) {
	data, _ := ioutil.ReadFile(file)
	signature := crypto.Sign(data, privateKey)
	writeSignatureFile(signature, file)
}

func VerifyFile(publicKey ed25519.PublicKey, file string) (verification bool) {
	data, _ := ioutil.ReadFile(file)
	signature := readSignatureFile(file)
	verification = crypto.Verify(data, signature, publicKey)
	return verification
}

func writeSignatureFile(signature []byte, file string) {
	ioutil.WriteFile(file+".str_s", signature, 0644)
}

func readSignatureFile(file string) (signature []byte) {
	signature, _ = ioutil.ReadFile(file + ".str_s")
	return signature
}
