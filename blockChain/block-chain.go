package blockChain

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
)

//桶名，这个桶用来装区块的信息
var BUCKET_NAME ="blocks"
//表示最新区块的key名
var LAST_KEY ="lastHash"
//存储区块数据的文件
var CHAINDB ="chain.db"

//区块链的结构体对象
type BlockChain struct {
	LastHash []byte
	BoltDb   *bolt.DB
}

//构建一条区块链，并返回一个区块链实例
func NewBlockChain() BlockChain{
	//打开存储区块数据的文件
	db,err:=bolt.Open(CHAINDB,0600,nil)
	if err != nil {
		panic(err.Error())
	}
	var bl BlockChain
	//检查创世区块是否已存在
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			bucket,err=tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash:=bucket.Get([]byte(LAST_KEY))
		if len(lastHash)==0{//没有创世区块
			//创建创世区块
			genesis :=CreatGenesisBlock()
			fmt.Printf("第一个区块的hash值:%x\n",genesis.Hash)
			//创建一个存储区块数据的文件
			bl =BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes,_:=genesis.Serialize()
			bucket.Put(genesis.Hash,genesisBytes)
			bucket.Put([]byte(LAST_KEY),genesis.Hash)
		} else{//有创世区块
			lastHash:=bucket.Get([]byte(LAST_KEY))
			lastBlockBytes:=bucket.Get(lastHash)
			lastBlock,err:=DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl=BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil
	})
	return bl
}

//构建一个BlockChain的方法，该方法可以将一个新生成的区块保存到chain.db文件中
func (bc BlockChain) SaveData(data []byte) (Block,error){
	db:=bc.BoltDb
	var er error
	var lastBlock *Block
	//先查询chain.db文件中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			er = errors.New("boltDb未创建，请重试！")
			return er
		}
		lastBlockBytes:=bucket.Get(bc.LastHash)
		lastBlock,_=DeSerialize(lastBlockBytes)
		return nil
	})

	//生成一个区块，把data存入新的区块中
	newBlock:=NewBlock(lastBlock.Height+1,data,lastBlock.Hash)

	//更新chain.db文件 把新生成的区块存入boltDb中
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		//key=hash,value=block's byte
		//区块的序列化
		newBlockBytes,_:=newBlock.Serialize()
		//把区块信息存到BoltDb中
		bucket.Put(newBlock.Hash,newBlockBytes)
		//更新代表最后一个区块hash值得记录
		bucket.Put([]byte(LAST_KEY),newBlock.Hash)
		return nil

	})
	return newBlock,er
}

//查询所有的区块信息，返回一个区块切片
func (bc BlockChain) QueryAllBlock() []*Block{
	blocks:=make([]*Block,0)
	db:=bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			panic("查询数据错误")
		}
		eachKey:=bc.LastHash
		prevHashBig:=new(big.Int)
		zeroBig:=big.NewInt(0)
		for {
			eachBlockBytes:=bucket.Get(eachKey)
			//反序列化后得到每一个区块
			eachBlock,_:=DeSerialize(eachBlockBytes)
			//将遍历的每一个区块结构体指针放到[]byte容器中
			blocks=append(blocks,eachBlock)
			prevHashBig.SetBytes(eachBlock.PrevHash)
			if prevHashBig.Cmp(zeroBig)==0 {
				break
			}
			eachKey=eachBlock.PrevHash
		}
		return nil
	})
	return blocks
}

//通过区块的高度查询某个具体的区块，返回区块实例
func (bc BlockChain) QueryBlockByHeight(height int64)*Block{
	if height<0{//目标高度小于0，参数不合法
		return nil
	}
	var block *Block
	db:=bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil{
			panic("查询数据失败")
		}
		hashKey:=bc.LastHash
		for {
			lastBlockBytes:=bucket.Get(hashKey)
			eachBlock,_:=DeSerialize(lastBlockBytes)
			if eachBlock.Height<height {//当给的数字超出区块链中区块的高度，直接返回
				break
			}
			if eachBlock.Height==height {
				block=eachBlock
				break
			}
			//遍历当前的区块高度与目标高度不一致，继续向前遍历
			//以eachBlock.PrevHash为Key,使用Get获取上一个区块的数据
			hashKey=eachBlock.PrevHash
		}
		return nil
	})
	return block
}