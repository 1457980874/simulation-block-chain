package blockChain

import (
	"bytes"
	"encoding/gob"
	"time"
)

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
//区块的序列化
func (bk Block)Serialize()([]byte,error){
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(bk)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}
//区块的反序列化
func DeSerialize(data []byte)(*Block,error){
	var block Block
	err:=gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err != nil {
		return nil,err
	}
	return &block,nil
}