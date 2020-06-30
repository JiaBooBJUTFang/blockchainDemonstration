package controllers



type Blockchain struct {
	chain []*Block
}

func Newblockchain() *Blockchain {
	block := NewGenesisBlock()
	return &Blockchain{chain: []*Block{block}}
}
func (bc *Blockchain) Addblock_To_blockchain(data string) {
	//bc.chain[len(bc.chain)-1].Hash
	prehash := []byte{}
	block := NewBlock(data, prehash)
	block.Block_high=bc.chain[len(bc.chain)-1].Block_high+1
	bc.chain = append(bc.chain, block)

}