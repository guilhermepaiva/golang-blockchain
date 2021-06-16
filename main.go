package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

/*A method that allows to create a hash based on the previous hash.*/
func (block *Block) DeriveHash() {
	info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	hash := sha256.Sum256(info)

	block.Hash = hash[:]
}

/*A function that creates the actual block*/
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()

	return block
}

/*A method that allow us to add a block to the chain*/
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

/*Let's create a function that creates the first block of the chain*/
func CreateFirstBlock() *Block {
	return CreateBlock("First Block", []byte{})
}

/*Build our initial blockchain using our first block*/
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{CreateFirstBlock()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("first block after genesis of the blockchain")
	chain.AddBlock("second block after genesis of the blockchain")
	chain.AddBlock("third block after genesis of the blockchain")

	for _, block := range chain.blocks {
		fmt.Println("*************************************************")
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("*************************************************")
		fmt.Println()
	}
}
