package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/aybabtme/color/brush"
)

func main() {
	fAll := flag.Bool("a", false, "toggle print all file")
	flag.Parse()

	var dirStr string
	if len(os.Args) > 1 {
		dirStr = os.Args[1]
	} else {
		dirStr = "./"
	}

	dir, err := os.Open(dirStr)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
	fileinfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	for _, v := range fileinfos {
		if !*fAll && strings.HasPrefix(v.Name(), ".") {
			continue
		}
		if v.IsDir() {
			fmt.Printf("%s \n", brush.DarkYellow(v.Name()))
			continue
		}
		if strings.Contains(v.Name(), ".go") {
			fmt.Printf("%s \n", brush.DarkGreen(v.Name()))
			continue
		}
		fmt.Printf("%s \n", v.Name())
	}
}
