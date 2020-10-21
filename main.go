package main

import (
	"DataCertProject/blockchain"
	"fmt"
)

func main() {
	block := blockchain.NewBlock(0, []byte{}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)
	fmt.Printf("区块Hash值：%x\n", block.Hash)
	fmt.Printf("区块的nonce值:%d\n", block.Nonce)
}
