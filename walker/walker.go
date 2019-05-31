package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// assertDir: 引数が有効なディレクトリか検証する
func assertDir(path string) {
	info, err := os.Stat(path)
	if !info.IsDir() {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
}

//Walk :指定されたディレクトリを再帰的に操作し、見つかったファイルを chan で返す。
func Walk(rootPath string) chan string {
	assertDir(rootPath)

	ch := make(chan string)
	go func() {
		err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(fmt.Errorf("prevent panic by handling failure accessing a path %q: %v", rootPath, err))
			}
			if !info.IsDir() {
				ch <- path
			}
			return nil
		})
		defer close(ch)
		if err != nil {
			log.Fatal(fmt.Errorf("error walking the path %q: %v", rootPath, err))
		}
	}()
	return ch
}
