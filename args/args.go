package args

import (
	"flag"
	"fmt"
	"log"
)

// Args struct
type Args struct {
	FileTypeFrom   string
	FileTypeTo     string
	RootFolderName []string
}

// ParseArgs : constructor of struct "args"
func ParseArgs() *Args {
	// Todo: file type check
	arg1 := flag.String("from", "jpg", "original file type to convert")
	arg2 := flag.String("to", "png", "file type you want to convert")

	flag.Parse()

	folder := flag.Args()
	if folder == nil {
		log.Fatal("have to chose directory to convert")
	}
	fmt.Println(folder)
	newArgs := &Args{
		FileTypeFrom:   *arg1,
		FileTypeTo:     *arg2,
		RootFolderName: flag.Args(),
	}

	return newArgs
}
