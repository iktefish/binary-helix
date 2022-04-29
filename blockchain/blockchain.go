package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash    []byte
	C_Id    []byte
	Ip_Port []byte
	OldHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.C_Id, b.Ip_Port, b.OldHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(cId string, ip_port string, oldHash []byte) *Block {
	block := &Block{[]byte{}, []byte(cId), []byte(ip_port), oldHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(cId string, ip_port string) {
	oldBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(cId, ip_port, oldBlock.OldHash)
	chain.Blocks = append(chain.Blocks, new)
}

func Origin() *Block {
	return CreateBlock("Origin", "0.0.0.0:0000", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Origin()}}
}

// func Exe() {
// 	CHAIN := InitBlockChain()
//
// 	for i, block := range CHAIN.Blocks {
// 		fmt.Println("\nORIGIN BLOCK:\t\n")
// 		fmt.Printf("\tblock.OldHash\t%v ~~>\t%x\n", i, block.OldHash)
// 		fmt.Printf("\tblock.C_Id\t%v ~~>\t%x\n", i, block.C_Id)
// 		fmt.Printf("\tblock.Ip_Port\t%v ~~>\t%x\n", i, block.Ip_Port)
// 		fmt.Printf("\tblock.Hash\t%v ~~>\t%x\n", i, block.Hash)
// 		fmt.Println("\n")
// 	}
//
// 	// chain.AddBlock("First Block")
// 	// chain.AddBlock("Second Block")
// 	// chain.AddBlock("Third Block")
//
// 	// for i, block := range chain.Blocks {
// 	// 	fmt.Printf("\tblock.OldHash\t%v ~~>\t%x\n", i, block.OldHash)
// 	// 	fmt.Printf("\tblock.Data\t%v ~~>\t%x\n", i, block.Data)
// 	// 	fmt.Printf("\tblock.Hash\t%v ~~>\t%x\n", i, block.Hash)
// 	// }
// }
