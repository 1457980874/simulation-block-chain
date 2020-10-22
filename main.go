package main

import (
	"fmt"
	"simulation_block_chain/blockChain"
)

func main() {
	block := blockChain.NewBlock(0, []byte{}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)
	fmt.Printf("区块Hash值：%x\n", block.Hash)
	fmt.Printf("区块的nonce值:%d\n", block.Nonce)
}
