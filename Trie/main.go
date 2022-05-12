package main

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

}
