package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func parseArguments() int {
	var length int
	flag.IntVar(&length, "len", 5, "Word length.")
	flag.Parse()
	return length
}

func main() {
	cwd, _ := os.Getwd()
	log.Printf("pwd: %s", cwd)
	file, err := os.Open("resources/easy_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordLen := parseArguments()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == wordLen {
			fmt.Println(word)
		}
	}
}
