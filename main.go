package main

import (
	"fmt"

	"github.com/richardsplit/golang_blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second(2) Block after Genesis")
	chain.AddBlock("Next(3) Block after Genesis")

	for _, block := range chain.Blocks {
		//fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}
}
