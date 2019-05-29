package main

import (
	"fmt"
	"log"
	"myConverter/args"
	"myConverter/converter"
	"myConverter/walker"
	"os"
	"path/filepath"
	"strings"
)

func execute(filePath string) {
	ext := strings.ToLower(filepath.Ext(filePath))

	fmt.Println(ext, "start")
	switch ext {
	case ".jpeg", ".jpg":
		converter.J2P(filePath)
	default:
	}
}

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
		execute(filePath)
	}
}
