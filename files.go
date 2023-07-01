package main

type file struct {
	fileName      string
	fileExtension string
}

// Returns File struct
func getNewFile(fileName, fileExtension string) *file {
	return &file{
		fileName:      fileName,
		fileExtension: fileExtension,
	}
}
