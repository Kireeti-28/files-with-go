package main

import (
	"fmt"
	"os"
)

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
