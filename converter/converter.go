package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"strings"
)

type imageFile struct {
	filePath     string
	fileBaseName string
	fileExt      string
}

func newImageFile(filePath string) *imageFile {
	newFile := &imageFile{}

	return newFile
}

func dec(filePath string, decodeType string) (image.Image, error) {
	reader, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	switch decodeType {
	case ".gif":
		return gif.Decode(reader)
	case ".jpeg", ".jpg":
		return jpeg.Decode(reader)
	case ".png":
		return png.Decode(reader)
	}
	return nil, nil
}

func enc(filePath string, encodeType string, m image.Image) error {
	f, err := os.Create(strings.TrimSuffix(filePath, path.Ext(filePath)) + "." + encodeType)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	switch encodeType {
	case ".gif":
		return gif.Encode(f, m, nil)
	case ".jpeg", ".jpg":
		return jpeg.Encode(f, m, nil)
	case ".png":
		return png.Encode(f, m)
	}
	return nil
}

//Convert : convert Jpeg file to PNG
func Convert(filePath string, decodeType string, encodeType string) {

	fmt.Println("Converting", filePath)

	m, err := dec(filePath, decodeType)
	if err != nil {
		log.Fatal(err)
	}

	err = enc(filePath, encodeType, m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Converted", filePath)
}
