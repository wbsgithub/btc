package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  uint
}

func main() {
	var xiaoming Person
	xiaoming.Name = "小明"
	xiaoming.Age = 10
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoming)
	if err != nil {
		log.Panic("编码出错")
	}
	fmt.Printf("编码后：%v\n", buffer.Bytes())
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	var daming Person
	err = decoder.Decode(&daming)
	if err != nil {
		log.Panic("解码出错")
	}
	fmt.Printf("解码后：%v\n", &daming)

}
