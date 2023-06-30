package main

import (
	"fmt"
	"log"
	"os"
)

type file struct {
	fileName      string
	fileExtension string
}

// Returns File struct
func getNewFile(fileName, fileExtension string) file {
	return file{
		fileName:      fileName,
		fileExtension: fileExtension,
	}
}

/**
 * notes:
 * os package had func to create file
 * method signature  od Create func Create(name string) (*File, error)
 */
func (f *file) createFile() error {
	file, err := os.Create(f.fileName + "." + f.fileExtension) // if sample.txt already existed seems like its replacing existed file with new file(erases data in existed file).
	fmt.Println(file)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Let's learn working with files in Go.")
	fmt.Println("=====================================")
	fmt.Println("Create File")

	f := getNewFile("sample", "txt") //  we can also get file name, extension via CLI. [Programatically 😉]
	err := f.createFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully created file")
	fmt.Println("=====================================")
}
