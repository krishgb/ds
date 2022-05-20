package trie

import (
	"fmt"
	"sort"
)

type node struct {
	children  map[byte]*node
	frequency int
}

type Trie struct {
	root *node
}

// initTrie initializes a node
func initTrie() *node {
	return &node{
		children:  map[byte]*node{},
		frequency: 0,
	}
}

// Insert(word)*Trie - inserts a word to the Trie
func (t *Trie) Insert(word string) *Trie {
	if t.root == nil {
		t.root = initTrie()
	}
	root := t.root
	for _, w := range word {
		if _, ok := root.children[b(w)]; !ok {
			root.children[b(w)] = initTrie()
		}
		root = root.children[b(w)]
	}
	root.frequency++
	return t
}

// Search(string)bool  - searches the given word in Trie
func (t *Trie) Search(word string) bool {
	if t == nil {
		return false
	}
	root := t.root
	for _, w := range word {
		if _, ok := root.children[b(w)]; !ok {
			return false
		}
		root = root.children[b(w)]
	}
	return root.frequency > 0
}

type str struct {
	word      string
	frequency int
}

// AutoComplete(word)([]string, error) - completes the given word
func (t *Trie) AutoComplete(word string) ([]str, error) {
	result := []str{}
	if t.root == nil {
		return result, fmt.Errorf("trie is empty")
	}

	root := t.root

	for _, w := range word {
		if _, ok := root.children[b(w)]; !ok {
			return result, fmt.Errorf(word, "not in trie")
		}
		root = root.children[b(w)]
	}

	var autoCompleteHelper func(string, *node, *[]str)
	autoCompleteHelper = func(word string, root *node, strs *[]str) {
		if root.frequency > 0 {
			*strs = append(*strs, str{word: word, frequency: root.frequency})
		}
		for _, w := range byteSort(root.children) {
			autoCompleteHelper(word+string(w), root.children[w], strs)
		}
	}
	autoCompleteHelper(word, root, &result)

	return result, nil
}

// Print() - print all the words in an asc order
func (t *Trie) Print() {
	if t.root == nil {
		return
	}
	root := t.root.children
	for _, w := range byteSort(root) {
		words, _ := t.AutoComplete(string(w))
		fmt.Println(words, "\v")
	}
}

func (t *Trie) Remove(word string) error {
	if word == "" || !t.Search(word) {
		return fmt.Errorf("\"%s\" not in trie", word)
	}

	remove := removeHelper(word, t.root.children[word[0]], 0)
	if remove {
		delete(t.root.children, word[0])
	}

	return nil
}

func removeHelper(word string, root *node, idx int) bool {

	if idx == len(word)-1 {
		if root.frequency > 1 {
			root.frequency--
			return false
		}
		return root.frequency == 1
	}

	if removeHelper(word, root.children[word[idx+1]], idx+1) {
		delete(root.children, word[idx+1])
		subNodes := 0
		for range root.children {
			subNodes++
		}
		return subNodes == 0
	}
	return false
}

func (t *Trie) Update(oldWord string, newWord string) bool {
	err := t.Remove(oldWord)
	if err != nil {
		return false
	}
	t.Insert(newWord)
	return true
}

// -----------------------------------Helper Functions
var b = func(a rune) byte {
	return byte(a)
}

var byteSort = func(a map[byte]*node) []byte {
	b := []byte{}
	for w := range a {
		b = append(b, w)
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return b
}
