package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hash_test(txt string) string {
	var buf string
	buf = txt
	hash := sha256.Sum256([]byte(buf))
	hashstr := hex.EncodeToString(hash[:])
	fmt.Println(hashstr)
	return hashstr
}
