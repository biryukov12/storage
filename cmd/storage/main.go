package main

import (
	"fmt"
	"github.com/biryukov12/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loaded", file)

	fmt.Println("restored", restoredFile)
}
