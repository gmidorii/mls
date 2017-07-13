package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dirArg := os.Args[1]
	fmt.Println(dirArg)
	dir, err := os.Open(dirArg)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
	fileinfos, err := dir.Readdir(0)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	for _, v := range fileinfos {
		fmt.Println(v)
	}
}
