package core

import (
	"github.com/boltdb/bolt"
	"log"
)

//在blockchain中将所有的数据块保存到了数据库中
//如果我们需要去访问其中的区块只能去迭代查询
//BoltDB允许对一个bucket里面的所有key进行迭代,但是所有的key
//都以字节存储.当在数据库很大的情况下,全部装入数据库是不现实的,
//因此需要一个区块链迭代器
type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

//返回一个迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator{
	return &BlockchainIterator{bc.Tip,bc.DB}
}

//返回链中的下一个块
func (bi *BlockchainIterator) Next() *Block{
	var block *Block

	err := bi.db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(BlockBucket))
		//根据迭代器的当前hash获取当前块的字节
		ecodeBlock := bucket.Get(bi.currentHash)
		//反序列化
		block = DeserializeBlock(ecodeBlock)
		//更新当前迭代器持有的hash
		//bi.currentHash = block.PreBlockHash
		return nil
	})
	if err != nil{
		log.Panic(err)
	}
	return block
}
















