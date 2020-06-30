package controllers

import (
	"crypto/sha256"
	"fmt"
	"log"

	"golang.org/x/crypto/ripemd160"
)

func ByteToString(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%d", b[i])
	}
	return s
}
func ByteToStringpub(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%x", b[i])
	}
	return s
}
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}


func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}


