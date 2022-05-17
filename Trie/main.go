package main

import (
	"bufio"
	"fmt"
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

	wordsToFind := []string{"computer", "golang", "programming", "hobby"}

	for i := 0; i < len(wordsToFind); i++ {

		found := trie.search(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}

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

func (t *trie) insert(word string) {
	// wordLength := len(word)
	current := t.root

	for _, letter := range word {
		//To get the index of a character,
		// For example c-a would be translated to (99–97)=2 which is index of c.
		index := letter - 'a'

		if current.childrens[index] == nil {
			current.childrens[index] = &node{character: letter}
		}

		current = current.childrens[index]
	}

	current.isWordEnd = true
}

func (t *trie) search(word string) bool {
	wordLength := len(word)
	current := t.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'

		if current == nil || current.childrens[index] == nil {
			return false
		}

		current = current.childrens[index]
	}

	if current.isWordEnd {
		return true
	}

	return false
}

func (t *trie) autoComplete(word string) {
	var words []string

	current := t.root

	for _, char := range word {

		index := char - 'a'

		if current.childrens[index] == nil {
			return
		}

		current = current.childrens[index]
	}

	t.root = current

	suggestions := t.traversal(word, words)

	fmt.Printf("The Auto Completion words for %s are :\n ", word)
	for _, suggestion := range suggestions {
		fmt.Println(suggestion)
	}

	return
}

func (t *trie) traversal(word string, words []string) []string {
	current := t.root

	if current.isWordEnd {
		words = append(words, word)
		return words
	}

	for _, localCurrent := range current.childrens {
		if localCurrent != nil {
			makeWord := word + string(localCurrent.character)
			index := localCurrent.character - 'a'
			t.root = current.childrens[index]
			words = t.traversal(makeWord, words)
		}
	}

	return words
}
