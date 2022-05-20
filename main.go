package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/krishgb/ds/trie"
)

func read(filename string) []string {
	result := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), " ")...)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return result
}

func main() {
	var t trie.Trie
	words := read("./trie/data.txt")
	for _, w := range words {
		t.Insert(w)
	}
	t.Print()
}
