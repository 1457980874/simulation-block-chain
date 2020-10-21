package blockChain

import "time"

type Block struct {
	Height int64
	TimeStamp int64
	Hash []byte
	Data []byte
	PrevHash []byte
	Version string
	Nonce int64
}

func NewBlock(height int64, data []byte, prevsHash []byte)(Block){
	block:=Block{
		Height:height,
		TimeStamp:time.Now().UnixNano(),
		Data:data,
		PrevHash:prevsHash,
		Version:"0x01",
	}
	pow:=NewPow(block)
	blockHash,nonce:=pow.run()

	block.Nonce=nonce
	block.Hash=blockHash

	return block
}
