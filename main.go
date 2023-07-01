package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Let's learn working with files in Go.")
	fmt.Println("=====================================")
	fmt.Println("Write File")

	f := initCreateFile()

	f.writeFile(getFileData())

	fmt.Println("Successfully writed file")
	fmt.Println("=====================================")
}

func getFileData() []byte {
	// writing json into file.
	type student struct {
		Id      int      `json:"id"`
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Hobbies []string `json:"hobbies"`
	}
	studentJson := student{
		Id:      1,
		Name:    "Kireeti",
		Age:     23,
		Hobbies: []string{"Coding", "Talking", "Overthinking"},
	}

	dat, err := json.Marshal(studentJson)

	if err != nil {
		log.Fatal(err)
	}

	return dat
}

func initCreateFile() *file {
	f := getNewFile("sample", "json") //  we can also get file name, extension via CLI. [Programatically 😉]
	err := f.createFile()
	if err != nil {
		log.Fatal(err)
	}
	return f
}
