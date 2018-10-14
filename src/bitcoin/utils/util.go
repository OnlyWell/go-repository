package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

//int64转为[]byte
func IntToHex(number int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, number)
	if err != nil{
		log.Panic(err)
	}

	return buff.Bytes()
}
