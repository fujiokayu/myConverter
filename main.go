package main

import (
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

	walker.Walk(folder)
}
