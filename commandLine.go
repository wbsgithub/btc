package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	fmt.Println("cli AddBlock")
	cli.bc.AddBlock(data)
	fmt.Println("添加区块成功")
}

func (cli *CLI) PrintBlockChain() {
	iterator := cli.bc.NewBlockChainIterator()
	for {
		block := iterator.Next()
		if block != nil {
			fmt.Println("====================================================================================")
			fmt.Printf("版本号:%d\n", block.Version)
			fmt.Printf("前区块hash:%x\n", block.PrevHash)
			fmt.Printf("梅克尔根:%x\n", block.MerkelRoot)
			fmt.Printf("时间戳:%d\n", block.Timestamp)
			fmt.Printf("难度值:%d\n", block.Difficulty)
			fmt.Printf("随机数:%d\n", block.Nonce)
			fmt.Printf("当前区块hash:%x\n", block.Hash)
			fmt.Printf("数据:%s\n", block.Data)
		}
		if len(block.PrevHash) == 0 {
			fmt.Println("读取完毕")
			break
		}
	}
}
