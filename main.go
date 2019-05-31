package main

import (
	"log"
	"myConverter/args"
	"myConverter/converter"
	"myConverter/walker"
	"os"
	"path/filepath"
	"strings"
)

// フォルダ探索によって見つかったファイルの拡張子が、decodeType と一致した場合に converter を起動する。
func execute(filePath string, decodeType string, encodeType string) {
	ext := strings.ToLower(filepath.Ext(filePath))

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
		execute(filePath, args.DecodeType, args.EncodeType)
	}
}
