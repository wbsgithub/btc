package main

import (
	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	db                 *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain) NewBlockChainIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

func (itr *BlockChainIterator) Next() *Block {
	var block Block
	itr.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		currentBlock := bucket.Get(itr.currentHashPointer)
		block = Deserilize(currentBlock)
		itr.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}
