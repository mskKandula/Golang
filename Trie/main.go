package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	//ALBHABET_SIZE total characters in english alphabet
	ALBHABET_SIZE = 26
	fileName      = "words_alpha.txt"
)

type node struct {
	character rune
	childrens [ALBHABET_SIZE]*node
	isWordEnd bool
}

type trie struct {
	root *node
}

func initTrie() *trie {
	return &trie{
		root: &node{},
	}
}
func main() {
	trie := initTrie()

	trie.readFileAndInsert(fileName)
}

func (t *trie) readFileAndInsert(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)

	if err != nil {
		log.Fatalln("Error while opening the file: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// To read word by word
	// scanner.Split(bufio.ScanWords)

	// To read line by line
	for scanner.Scan() {
		lineText := scanner.Text()
		// insert after converting to lowercase
		t.insert(strings.ToLower(lineText))
	}

	if err = scanner.Err(); err != nil {
		log.Fatalln("Error while reading file: ", err)
	}
}

func (t *trie) insert(word string) {}
