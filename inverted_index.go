package KSE

import (
	"errors"
	"github.com/zbw0046/KSE/TRIE"
)

type PostingEntry struct{
	DocId int
	Location []int
}

type PostingList struct{
	Entries map[int]*PostingEntry // map from docId(int) to entry
}

type Header struct{
	Word string
	DocFreq int
	*PostingList
}

type InvertedIndex struct{
	HeadersTrie *TRIE.Trie
}

func NewInvertedIndex() *InvertedIndex{
	return &InvertedIndex{
		HeadersTrie: TRIE.NewTrie(),
	}
}

func NewPostingList() *PostingList{
	return &PostingList{
		Entries: make(map[int]*PostingEntry),
	}
}

func NewHeader(word string) *Header{
	return &Header{
		Word: word,
		PostingList: NewPostingList(),
		DocFreq: 0,
	}
}

func (h *InvertedIndex) AddWord(word string, docId int, loc []int) error {
	if h.HeadersTrie == nil{
		return errors.New("header trie has not been built yet!\n")
	}
	var header *Header
	if headerObject, err := h.HeadersTrie.GetObjectFromTrie(word); err == nil{
		header = headerObject.(*Header)
	} else {
		header = NewHeader(word)
		h.HeadersTrie.AddObjectToTrie(word, header)
	}


}

func (h *InvertedIndex) GetEntriesFromKeyword(keyword string) []*PostingEntry{
	return nil
}