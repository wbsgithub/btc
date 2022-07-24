package main

import (
	"github.com/boltdb/bolt"
	"log"
)

/**
定义区块链结构
*/
type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB
	tail []byte //存储最后一个区块的哈希
}

const blockChainDB = "blockchain.db"
const blockBucket = "blockBucket"

func NewBlockChain() *BlockChain {
	//最后一个区块的哈希，从数据库中读出来
	var lastHash []byte
	db, err := bolt.Open(blockChainDB, 0600, nil)
	if err != nil {
		log.Panic("")
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket()失败")
			}
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash, genesisBlock.Serilize())
			bucket.Put([]byte("lastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})

	return &BlockChain{
		db:   db,
		tail: lastHash,
	}
}

func GenesisBlock() *Block {
	return NewBlock("创世区块", nil)
}

func (bc *BlockChain) AddBlock(data string) {
	db := bc.db
	prevHash := bc.tail
	//lastBlock := blockChain.blocks[len(blockChain.blocks)-1]
	//prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	blockBytes := block.Serilize()
	//blockChain.blocks = append(blockChain.blocks, block)
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空")
		} else {
			bucket.Put(block.Hash, blockBytes)
			bucket.Put([]byte("lastHashKey"), block.Hash)
			bc.tail = block.Hash
		}
		return nil
	})
}
