package args

import (
	"flag"
	"log"
)

// Args struct
type Args struct {
	DecodeType     string
	EncodeType     string
	RootFolderName []string
}

// ParseArgs : constructor of struct "args"
func ParseArgs() *Args {
	arg1 := flag.String("from", "jpg", "original file type to convert")
	arg2 := flag.String("to", "png", "file type you want to convert")

	flag.Parse()

	// フォルダが指定されているかチェックする
	folder := flag.Args()
	if len(folder) == 0 {
		log.Fatal("have to choose directory to convert")
	}

	newArgs := &Args{
		DecodeType:     *arg1,
		EncodeType:     *arg2,
		RootFolderName: flag.Args(),
	}

	return newArgs
}
