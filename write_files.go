package main

import (
	"log"
	"os"
)

func (f *file) writeFile(textBytes []byte) {
	path := f.fileName + "." + f.fileExtension

	err := os.WriteFile(path, textBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
