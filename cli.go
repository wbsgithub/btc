package main

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA "add data to blockchain"
	printChain "print all blockchain data"
`

func (cli *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块")
		fmt.Println("参数长度：", len(args))
		fmt.Println("args[2]=='--data' :", args[2] == "--data")
		//确保命令有效
		if len(args) == 4 && args[2] == "--data" {
			cli.AddBlock(args[3])
		}
	case "printChain":
		cli.PrintBlockChain()
		fmt.Println("打印区块")
	default:
		fmt.Printf(Usage)
	}

}
