package encoding

import (
	"encoding/base64"

	"golang.org/x/crypto/ed25519"
)

func PublicKeyToEncoded(publicKey ed25519.PublicKey) (publicKeyEncoded string) {
	publicKeyEncoded = base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(publicKey)
	return publicKeyEncoded
}

func EncodedToPublicKey(publicKeyEncoded string) (publicKey ed25519.PublicKey) {
	publicKey, _ = base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(publicKeyEncoded)
	return publicKey
}
