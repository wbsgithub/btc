package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

/**
定义区块结构
*/
type Block struct {
	Version    uint64
	PrevHash   []byte
	MerkelRoot []byte
	Timestamp  uint64
	Difficulty uint64
	Nonce      uint64
	Hash       []byte
	Data       []byte
}

func Uint64ToByte(data uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, data)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

func NewBlock(data string, prevHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevHash,
		MerkelRoot: []byte{},
		Timestamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{}, //先填空，后计算 //TODO
		Data:       []byte(data),
	}
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

func (block *Block) toByte() []byte {
	return []byte{}
}

func (block *Block) SetHash() {
	//var blockInfo []byte
	//blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	//blockInfo = append(blockInfo, block.PrevHash...)
	//blockInfo = append(blockInfo, block.MerkelRoot...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Timestamp)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	//blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	//blockInfo = append(blockInfo, block.Data...)

	blockInfoByte := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.Timestamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockInfo := bytes.Join(blockInfoByte, []byte{})
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
