package args

import (
	"flag"
)

// Args struct
type Args struct {
	RootFolderName []string
}

// ParseArgs : constructor of struct "args"
func ParseArgs() *Args {

	flag.Parse()
	newArgs := &Args{
		RootFolderName: flag.Args(),
	}

	return newArgs
}
