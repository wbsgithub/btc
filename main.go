package main

import (
	"strconv"
)

func main() {
	chain := NewBlockChain()
	for i := 1; i < 10; i++ {
		chain.AddBlock(strconv.Itoa(i) + "次添加区块")
	}
	//for i, block := range chain.blocks {
	//	fmt.Printf("========================前区块高度:%d======================\n", i)
	//	fmt.Printf("前区块hash:%x\n", block.PrevHash)
	//	fmt.Printf("当前区块hash:%x\n", block.Hash)
	//	fmt.Printf("数据:%s\n", block.Data)
	//}

}
