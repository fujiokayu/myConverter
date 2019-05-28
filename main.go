package main

import (
	"fmt"
	"log"
	"myConverter/args"
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

	fmt.Println(folder, "is Directory!")
}
