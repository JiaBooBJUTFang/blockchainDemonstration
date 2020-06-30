// block
package controllers

import (
	"bytes"
	"os"

	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Block_high int64  `json:"blockhigh"` 
	Prevhash   []byte `json:"prehash"`  
	MerkelRoot []byte
	TimeStamp  int64  `json:"timestamp"` 
	//Bits       int    //难度值
	Nonce int64  `json:"nonce"`
	Hash  []byte `json:"hash"`
	Data  []byte `json:"data"`
}

func NewBlock(data string, prev []byte) *Block {
	var block Block

	block = Block{
		Block_high: 1,
		Prevhash:   prev,
		//MerkelRoot: []byte{},
		//Bits:       1,
		Nonce:     1,
		Data:      []byte(data),
		TimeStamp: time.Now().Unix()}

	(&block).sethash()
	//Showblock(&block)
	return &block
}
func NewBlock_simple(data string, blockhigh int64) *Block {
	var block Block
	block = Block{
		Block_high: blockhigh,
		Data:       []byte(data)}

	return &block
}
func (block *Block) set_simple_hash(time string) []byte {
	tem := [][]byte{
		InttoByte(block.Block_high),
		[]byte(time),
		block.Data,
	}
	data := bytes.Join(tem, []byte{})

	hash := sha256.Sum256(data)

	return hash[:]
}
func NewGenesisBlock() *Block {
	block := NewBlock("", []byte{})
	return block
}
func (block *Block) sethash() {
	tem := [][]byte{
		InttoByte(block.Block_high),
		block.Prevhash,
		block.MerkelRoot,
		[]byte(strconv.FormatInt(block.TimeStamp, 10)),
		InttoByte(block.Nonce),
		block.Data,
	}
	data := bytes.Join(tem, []byte{})

	hash := sha256.Sum256(data)

	block.Hash = hash[:]
}

func Showblock(b *Block) {
	fmt.Println("Block:#", b.Block_high)
	fmt.Println("TimeStamp :", time.Unix(b.TimeStamp, 0).Format("2006-01-02 03:04:05 PM"))
	fmt.Println("Nonce :", b.Nonce)
	fmt.Println("Prev :", b.Prevhash)
	fmt.Print("Hash :")
	for _, data := range b.Hash {
		fmt.Printf("%x ", data)
	}
	fmt.Println()
	fmt.Println("Data: ", string(b.Data))
	fmt.Println("------------------------------------------------------------------------------")

}
func InttoByte(data int64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, data)
	if err != nil {
		fmt.Println("InttoByte err=", err)
		os.Exit(0)
	}
	return buf.Bytes()
}
