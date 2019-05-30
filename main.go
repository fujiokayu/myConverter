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

func execute(filePath string, decodeType string, encodeType string) {

	ext := strings.ToLower(filepath.Ext(filePath))
	fmt.Println(ext, strings.ToLower("."+decodeType))
	if ext == strings.ToLower("."+decodeType) {
		converter.Convert(filePath, ext, encodeType)
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
		execute(filePath, args.FileTypeFrom, args.FileTypeTo)
	}
}
