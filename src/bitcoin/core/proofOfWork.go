package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"bitcoin/utils"
)

var (
	maxNonce = math.MaxInt64
	//hashcount = 0
)
const targetBits = 24
type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork{
	//将target初始化为大整数1
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

//准备数据
func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data := bytes.Join([][]byte{
		pow.block.PreBlockHash,
		pow.block.Data,
		utils.IntToHex(pow.block.TimeStamp),
		utils.IntToHex(int64(targetBits)),
		utils.IntToHex(int64(nonce)),
	},[]byte{})
	return data
}
//工作量证明
func (pow *ProofOfWork) Run() (int, []byte){
	var hashInt big.Int //hash的整形表示
	var hash [32]byte

	nonce := 0 //计数器
	fmt.Printf("Mining the block containing %s",pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		fmt.Printf("%x \n",hash)
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r %x \n",hash)
			//fmt.Printf("hash计算次数: %d",hashcount)
			break
		} else {
			nonce ++
			//hashcount ++
		}
	}
	fmt.Printf("\n\n")
	return nonce,hash[:]
}

//区块验证
func (pow *ProofOfWork) Validate() bool{
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid

}
