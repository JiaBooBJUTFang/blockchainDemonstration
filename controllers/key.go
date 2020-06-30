package controllers

import (

	"crypto/ecdsa"
)

const (
	version            = byte(0x00)
	addreddChechsumLen = 4
	privKeyBytesLen    = 32
)

type GKey struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

var publicKey []byte
var privateKey []byte

func (k GKey) GetPrivKey() []byte {
	d := k.PrivateKey.D.Bytes()
	b := make([]byte, 0, privKeyBytesLen)
	priKey := paddedAppend(privKeyBytesLen, b, d) // []bytes type
	// s := byteToString(priKey)
	return priKey
}

func (k GKey) GetPubKey() []byte {
	pubKey := append(k.PublicKey.X.Bytes(), k.PrivateKey.Y.Bytes()...) // []bytes type
	
	return pubKey
}