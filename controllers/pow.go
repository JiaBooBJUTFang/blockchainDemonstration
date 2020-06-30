package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 8

type ProofofWork struct {
	block       *Block
	target_hash *big.Int
}

//0001
//3 0 2 1 1 2 0 3
func NewProofofWork(b *Block) *ProofofWork {
	var target = big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := ProofofWork{
		block:       b,
		target_hash: target,
	}
	return &pow
}
func (pow *ProofofWork) predata(nonce int64) []byte {	
	block := pow.block
	tem := [][]byte{
		InttoByte(block.Block_high),
		InttoByte(nonce),
		block.Data,
	}
	data := bytes.Join(tem, []byte{})
	return data
}
func mine(data []byte) [32]byte {
	hash := sha256.Sum256(data)
	return hash
}
func (pow *ProofofWork) run() (int64, []byte) {
	fmt.Println("mining............")
	var nonce int64 = 0
	var hash_Int big.Int
	var hash [32]byte
	for nonce < math.MaxInt64 {
		data := pow.predata(nonce)
		hash = sha256.Sum256(data)
		if hash[0] == 0 && hash[1] == 0 {
			hash_Int.SetBytes(hash[:])
			if hash_Int.Cmp(pow.target_hash) < 0 {
				hex.EncodeToString(hash[:])
				break
			}
			nonce++
		} else {
			nonce++
		}


	}
	return nonce, hash[:]
}
