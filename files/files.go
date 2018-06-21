package files

import (
	"io/ioutil"

	"github.com/patrickmcnamara/soter/crypto"

	"golang.org/x/crypto/ed25519"
)

func WriteKeyfile(keyfile string, privateKey ed25519.PrivateKey) {
	privateKey = crypto.EncryptKey(privateKey)
	ioutil.WriteFile(keyfile, privateKey, 0600)
}

func ReadKeyfile(keyfile string) (privateKey ed25519.PrivateKey) {
	privateKey, _ = ioutil.ReadFile(keyfile)
	privateKey = crypto.DecryptKey(privateKey)
	return privateKey
}

func BackupKeyfile(file, keyfile string) {
	data, _ := ioutil.ReadFile(keyfile)
	ioutil.WriteFile(file, data, 0600)
}

func RestoreKeyfile(file, keyfile string) {
	data, _ := ioutil.ReadFile(file)
	ioutil.WriteFile(keyfile, data, 0600)
}

func SignFile(file string, privateKey ed25519.PrivateKey) {
	data, _ := ioutil.ReadFile(file)
	signature := crypto.Sign(data, privateKey)
	writeSignatureFile(signature, file)
}

func VerifyFile(file string, publicKey ed25519.PublicKey) (verification bool) {
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
