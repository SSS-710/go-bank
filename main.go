package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ Account created: %s (%s %s)\n", acc.Number, fname, lname)

	return acc
}

func seedAccounts(store Storage) {
	accounts := []struct {
		fname string
		lname string
		pw    string
	}{
		{"Anthony", "GG", "hunter88888"},
		{"John", "Doe", "password123"},
		{"Alice", "Smith", "secure456"},
	}

	for _, acc := range accounts {
		seedAccount(store, acc.fname, acc.lname, acc.pw)
	}
}

func main() {
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("🌱 Seeding database...")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	log.Println("🚀 API Server running on http://localhost:3000")
	server.Run()
}
