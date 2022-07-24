package main

func main() {
	chain := NewBlockChain()
	cli := CLI{chain}
	cli.Run()
	//for i := 1; i < 2; i++ {
	//	chain.AddBlock(strconv.Itoa(i) + "次添加区块")
	//}
	//iterator := chain.NewBlockChainIterator()
	//for {
	//	block := iterator.Next()
	//	if block != nil {
	//		fmt.Println("====================================================================================")
	//		fmt.Printf("前区块hash:%x\n", block.PrevHash)
	//		fmt.Printf("当前区块hash:%x\n", block.Hash)
	//		fmt.Printf("数据:%s\n", block.Data)
	//	}
	//	if len(block.PrevHash) == 0 {
	//		fmt.Println("读取完毕")
	//		break
	//	}
	//}
}
