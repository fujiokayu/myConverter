package converter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"strings"
)

//J2P : convert Jpeg file to PNG
func J2P(filePath string) {

	fmt.Println("Converting", filePath)

	reader, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, err := jpeg.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(strings.TrimSuffix(filePath, path.Ext(filePath)) + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, m); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Converted", filePath)
}
