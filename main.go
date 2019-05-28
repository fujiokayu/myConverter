package main

import (
	"fmt"
	"log"
	"myConverter/args"
	"myConverter/walker"
	"os"
	"strings"
)

func main() {
	args := args.ParseArgs()
	folder := strings.Join(args.RootFolderName, " ")

	info, err := os.Stat(folder)
	if !info.IsDir() {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	ch := walker.Walk(folder)
	for filePath := range ch {
		fmt.Println(filePath)
	}
}
