package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	TimeStamp int64
	Data []byte
	PreBlockHash []byte
	Hash []byte
	Nonce int
}

//func (b *Block) SetHash(){
//	timestamp := []byte(strconv.FormatInt(b.TimeStamp,10))
//	headers := bytes.Join([][]byte{b.Data,b.PreBlockHash,b.Hash}, []byte(timestamp))
//	hash := sha256.Sum256(headers)
//	b.Hash = hash[:]
//}

//创建新块
func NewBlock(data string, preHash []byte) *Block{
	newBlock := &Block{
		TimeStamp: time.Now().Unix(),
		Data:[]byte(data),
		PreBlockHash:preHash,
		Nonce:0,
	}
	//工作量证明
	pow := NewProofOfWork(newBlock)

	nonce, hash := pow.Run()
	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	return newBlock
}

//serializable
func (b *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil{
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化
func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}




