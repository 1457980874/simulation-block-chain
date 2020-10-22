package blockChain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"simulation_block_chain/until"
)

const DIFF  =10

type ProofOfWork struct {
	Target *big.Int
	Block Block
}

func NewPow(block Block)ProofOfWork{
	target:=big.NewInt(1)
	fmt.Println(target)
	target.Lsh(target,255-DIFF)
	fmt.Println(target)
	pow:=ProofOfWork{
		Target: target,
		Block:  block,
	}
	return pow
}

func (p ProofOfWork) run()([]byte,int64){
	var nonce int64
	bigBlock := new(big.Int)
	var block256Hash []byte
	for {
		block:=p.Block

		heightBytes,_:=until.IntToByte(block.Height)
		timeBytes,_:=until.IntToByte(block.TimeStamp)
		versionBytes:=until.StringToByte(block.Version)

		nonceBytes,_:=until.IntToByte(nonce)

		blockBytes:=bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})

		//fmt.Println("走没走啊！",blockBytes)
		sha256Hash:=sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash=sha256Hash.Sum(nil)
		//fmt.Println(block256Hash)
		//fmt.Println(block.Hash)
		fmt.Println(nonce)
		bigBlock=bigBlock.SetBytes(block256Hash)
		if p.Target.Cmp(bigBlock)==1 {
			break
		}
		nonce++
	}
	return block256Hash,nonce

}




