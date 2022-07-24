package main

import (
	"fmt"
	"os"
)

func main() {
	for i, cmd := range os.Args {
		fmt.Printf("args[%d]:%s\n", i, cmd)
	}

}
