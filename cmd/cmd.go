package cmd

import (
	"github.com/patrickmcnamara/soter/crypto"
	"github.com/patrickmcnamara/soter/encoding"
	"github.com/patrickmcnamara/soter/files"
)

func GenerateKeypair(keyfile string) (publicKeyEncoded string) {
	publicKey, privateKey := crypto.GenerateKeypair()
	publicKeyEncoded = encoding.PublicKeyToEncoded(publicKey)
	files.WriteKeyfile(keyfile, privateKey)
	return publicKeyEncoded
}

func BackupKeyfile(backupKeyfile, keyfile string) {
	files.BackupKeyfile(backupKeyfile, keyfile)
}

func RestoreKeyfile(backupKeyfile, keyfile string) {
	files.RestoreKeyfile(backupKeyfile, keyfile)
}

func ChangeKeyfilePassword(keyfile string) {
	files.ChangeKeyfilePassword(keyfile)
}

func Sign(keyfile, file string) {
	privateKey := files.ReadKeyfile(keyfile)
	files.SignFile(privateKey, file)
}

func Verify(publicKeyEncoded, file string) (verification bool) {
	publicKey := encoding.EncodedToPublicKey(publicKeyEncoded)
	verification = files.VerifyFile(publicKey, file)
	return verification
}

func PrintPublicKey(keyfile string) (publicKeyEncoded string) {
	privateKey := files.ReadKeyfile(keyfile)
	publicKey := crypto.GetPublicKey(privateKey)
	publicKeyEncoded = encoding.PublicKeyToEncoded(publicKey)
	return publicKeyEncoded
}
