package blockChain

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"simulation block chain/until"
)

const DIFF  =20

type ProofOfWork struct {
	Target *big.Int
	Block Block
}

func NewPow(block Block)ProofOfWork{
	target:=big.NewInt(1)
	target.Lsh(target,255-DIFF)
	pow:=ProofOfWork{
		Target: target,
		Block:  block,
	}
	return pow
}

func (p ProofOfWork)run()([]byte,int64){
	var nonce int64
	bigBlock := new(big.Int)
	var block256Hash []byte
	for {
		block:=p.Block

		heightBytes,_:=until.IntToByte(block.Height)
		timeBytes,_:=until.IntToByte(block.TimeStamp)
		versionBytes:=until.StringToByte(block.Version)

		nonceBytes,_:=until.IntToByte(block.Nonce)

		blockBytes:=bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})
		sha256Hash:=sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash=sha256Hash.Sum(nil)

		bigBlock=bigBlock.SetBytes(block256Hash)
		if p.Target.Cmp(bigBlock)==1 {
			break
		}
		nonce++
	}
	return block256Hash,nonce

}




