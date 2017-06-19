package main

import (
	"fmt"
	"log"
	"github.com/yjhmelody/redgo"
)

func main() {
	conn, err := redgo.Dial("127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Excute("set", "age", "20")
	if err != nil {
		log.Fatal(err)
	}
	raw, err := conn.ReadRaw()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(raw)
}
