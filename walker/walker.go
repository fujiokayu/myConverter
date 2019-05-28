package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//Walk :
func Walk(rootPath string) chan string {
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
