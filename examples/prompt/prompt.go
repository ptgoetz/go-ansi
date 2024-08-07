package main

import (
	"fmt"
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"log"
	"os"
)

func main() {
	console := ansi.NewConsole(os.Stdout)
	var s string
	_, err := console.Prompt("What's your name?: ", &s)
	if err != nil {
		log.Fatal(err)
	}
	var age int
	_, err = console.Prompt("What's your age?: ", &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s. You are %d years old.\n", s, age)
}
