package main

import (
	bit "./core"
)

func main() {
	prototype()
}

func prototype() {
	bc := bit.NewBlockchain()
	bc.AddBlock("Send 1 BTC to harry")
	bc.AddBlock("Send 1 BTC to rose")
	//for _,block := range bit.Blocks{
	//	fmt.Printf("hash: %x \n",block.Hash)
	//	fmt.Printf("preHash: %x \n",block.PreBlockHash)
	//	fmt.Printf("data: %s \n",block.Data)
	//	pow := bit.NewProofOfWork(block)
	//	fmt.Printf("Pow: %s \n",strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//}
}
