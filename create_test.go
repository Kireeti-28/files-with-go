package main

import (
	"log"
	"testing"
)

func TestCreateFiles(t *testing.T) {
	cases := []struct{
		fileName string
		fileExtension string
	}{
		{
			fileName: "sample",
			fileExtension: "txt",
		},
		{
			fileName: "database",
			fileExtension: "json",
		},
	}

	for i := 0; i < len(cases); i++ {
		caseT := cases[i]
		f := getNewFile(caseT.fileName, caseT.fileExtension)
		err := f.createFile()
		if err != nil {
			log.Fatal(err)
		}
	}
}

