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
	fLine := flag.Bool("l", false, "print one line")
	flag.Parse()

	var dirStr string
	if size := len(os.Args); size > 1 && !strings.Contains(os.Args[size-1], "-") {
		dirStr = os.Args[size]
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

	for i := 0; i < len(fileinfos); i += 3 {
		var index int
		if i >= len(fileinfos)-2 {
			index = len(fileinfos)
		} else {
			index = i + 3
		}
		for j := i; j < index; j++ {
			if !print(fileinfos[j], *fAll) {
				continue
			}
			if *fLine {
				fmt.Printf("\n")
			}
		}
		if !*fLine {
			fmt.Printf("\n")
		}
	}
}

func print(file os.FileInfo, fAll bool) bool {
	if !fAll && strings.HasPrefix(file.Name(), ".") {
		return false
	}
	if file.IsDir() {
		fmt.Printf("%s\t", brush.DarkYellow(file.Name()))
		return true
	}
	if strings.Contains(file.Name(), ".go") {
		fmt.Printf("%s\t", brush.DarkGreen(file.Name()))
		return true
	}
	fmt.Printf("%s\t", file.Name())
	return true
}
