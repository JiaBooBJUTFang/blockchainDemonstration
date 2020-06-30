package controllers

type Distributed_Blockchain struct {
	Peer1 *Blockchain
	Peer2 *Blockchain
	Peer3 *Blockchain
}

func NewDistributedBlockchain() *Distributed_Blockchain {
	dis_chain := Distributed_Blockchain{
		Peer1: Newblockchain(),
		Peer2: Newblockchain(),
		Peer3: Newblockchain()}
	return &dis_chain
}
func (db *Distributed_Blockchain) InitDistributeChain() {
	p1 := db.Peer1
	p2 := db.Peer2
	p3 := db.Peer3

	for i := 0; i < N; i++ {
		p1.Addblock_To_blockchain("")
		p2.Addblock_To_blockchain("")
		p3.Addblock_To_blockchain("")
	}
}
