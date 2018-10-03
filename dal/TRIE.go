package dal

import (
	"errors"
)

type TrieNode struct{
	char byte // the char of this node
	data interface{} // the data corresponding to the word (e.g. the header of the posting list)
	next map[byte]*TrieNode // the link to the longer word
}

type Trie struct{
	root *TrieNode
}

func NewTrie() *Trie{
	rootNode := NewTrieNode(byte('r'))
	return &Trie{rootNode}
}

func NewTrieNode(char byte) *TrieNode{
	return &TrieNode{
		char: char,
		next: make(map[byte]*TrieNode),
		data: nil,
	}
}

func (t *Trie) AddObjectToTrie(word string, content interface{}) {
	pNode := t.root
	idx := 0
	for idx < len(word){
		char := byte(word[idx])
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

func (t *Trie) GetObjectFromTrie(word string) (interface{}, error) {
	pNode := t.root
	idx := 0
	for idx < len(word){
		char := byte(word[idx])
		var ok bool
		if pNode, ok = pNode.next[char]; !ok{
			return nil, errors.New("No this object\n")
		}
		idx++
	}
	return pNode.data, nil
}