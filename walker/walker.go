package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//Walk :
func Walk(rootPath string) {
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(fmt.Errorf("error walking the path %q: %v", rootPath, err))
	}
}
