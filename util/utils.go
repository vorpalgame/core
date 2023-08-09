package util

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func GetFile(fileName string) *os.File {
	log.Println(fileName)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()
	return f
}

func LoadImage(imageFileName string) *image.Image {
	f, err := os.Open(imageFileName)
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal("Error during decode: ", err)
	}
	return &img
}

func Write(buffer *image.RGBA) {
	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(out, buffer)
}
