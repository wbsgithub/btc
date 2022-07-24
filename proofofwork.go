package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block *Block
	//一个非常大的数，它又丰富的方法；比较，赋值
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	tmp := big.Int{}
	tmp.SetString(targetStr, 16)
	pow.target = &tmp
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	block := pow.block
	var nonce uint64
	var hash [32]byte
	for {
		blockInfoByte := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.Timestamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(blockInfoByte, []byte{})
		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			//找到了
			fmt.Printf("挖矿成功！hash： %x，nonce:%d\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}
}
