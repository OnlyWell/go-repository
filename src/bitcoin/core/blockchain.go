package core

import (
	"github.com/boltdb/bolt"
	"log"
)

const (
	DbFile = "blockchain"
	BlockBucket = "blockchain.db"
)
type Blockchain struct {
	//Blocks []*Block
	Tip []byte //用于存储最后一个区块的hash
	DB *bolt.DB //数据库
}

//新增一个块
//TODO
//func (bc *Blockchain) AddBlock(data string){
//	preHash := bc.Tip
//	newBlock := NewBlock(data,preHash)
//
//	bc.Blocks = append(bc.Blocks,newBlock)
//}

//增加一个新的区块
func (bc *Blockchain) AddBlock(data string){
	var lastHash []byte
	//查询最后一个块的hash
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil{
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucket))
		//将当前块持久化
		err1 := b.Put(newBlock.Hash,newBlock.Serialize())
		//更新数据库中的最新块的hash
		err1 = b.Put([]byte("l"),newBlock.Hash)
		if err1 != nil{
			log.Panic(err1)
		}
		//设置区块链的tip
		bc.Tip = newBlock.Hash
		return nil
	})
}
//创建创世块
func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block", []byte{})
}

//func NewBlockchain() *Blockchain{
//	return &Blockchain{[]*Block{NewGenesisBlock()}}
//}
/*
	持久化
 */
func NewBlockchain() *Blockchain{
	var tip []byte
	db,err := bolt.Open(DbFile,0600,nil)

	err = db.Update(func(tx *bolt.Tx) error {
		//获取数据表对象
		b := tx.Bucket([]byte(BlockBucket))
		//如果数据表对象不存在
		if b == nil{
			//初始化创世块
			genesis := NewGenesisBlock()
			//构建新的数据表
			b,err := tx.CreateBucket([]byte(BlockBucket))
			if err != nil{
				log.Panic(err)
			}
			//将创世区块的存入表,块hash作为key,块的序列化值作为value
			err = b.Put(genesis.Hash, genesis.Serialize())
			//用l保存创世块的hash
			err = b.Put([]byte("l"), genesis.Hash)
			//区块链的tip保存最后一个hash
			tip = genesis.Hash
		} else {
			//如果存在表,则取出最后一个hash,用于构建区块链对象
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil{
		log.Panic()
	}

	bc := Blockchain{tip,db}

	return &bc
}
