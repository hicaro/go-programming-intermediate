package main

import (
	"fmt"
	"log"
	"os"
)

func logging() {
	into_file()
}

func spits() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}

func into_file() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	spits()
}
