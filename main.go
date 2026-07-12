package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%+v\n", store)
	// server := NewAPIServer(":3000", store)
	// server.Run()
}
