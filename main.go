package main

import (
	"fmt"
	"strconv"

	"github.com/synycboom/simple-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
