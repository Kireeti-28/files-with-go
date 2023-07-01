package main

import (
	"fmt"
	"log"
)

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
