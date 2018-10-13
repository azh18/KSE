package dal

import (
	"errors"
)

type Char uint8

type TrieNode struct{
	char Char // the char of this node
	data interface{} // the data corresponding to the word (e.g. the header of the posting list)
	next map[Char]*TrieNode // the link to the longer word
}

type Trie struct{
	root *TrieNode
}

func NewTrie() *Trie{
	rootNode := NewTrieNode('r')
	return &Trie{rootNode}
}

func NewTrieNode(char Char) *TrieNode{
	return &TrieNode{
		char: char,
		next: make(map[Char]*TrieNode),
		data: nil,
	}
}

// Add an element into the Trie
func (t *Trie) AddObjectToTrie(word string, content interface{}) {
	pNode := t.root
	idx := 0
	for idx < len(word){
		char := Char(word[idx])
		if node, ok := pNode.next[char]; ok{
			pNode = node
		} else {
			pNode.next[char] = NewTrieNode(char)
			pNode = pNode.next[char]
		}
		idx++
	}
	pNode.data = content
}

// Get the elements of leaves according to the word
func (t *Trie) GetObjectFromTrie(word string) (interface{}, error) {
	pNode := t.root
	idx := 0
	for idx < len(word){
		char := Char(word[idx])
		var ok bool
		if pNode, ok = pNode.next[char]; !ok{
			return nil, errors.New("No this object\n")
		}
		idx++
	}
	if pNode.data == nil{
		return nil, errors.New("No data in this node.\n")
	}
	return pNode.data, nil
}