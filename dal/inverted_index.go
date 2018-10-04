package dal

import (
	"errors"
	"strconv"
	"strings"
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
	HeadersTrie *Trie
}

func NewInvertedIndex() *InvertedIndex{
	return &InvertedIndex{
		HeadersTrie: NewTrie(),
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

// call by only once for each (doc, word) pair
func (h *InvertedIndex) AddWords(docId int, wordLoc WordLoc) error {
	if h.HeadersTrie == nil{
		return errors.New("header trie has not been built yet!\n")
	}
	var header *Header
	for word, loc := range wordLoc {
		if headerObject, err := h.HeadersTrie.GetObjectFromTrie(word); err == nil{
			header = headerObject.(*Header)
		} else {
			header = NewHeader(word)
			h.HeadersTrie.AddObjectToTrie(word, header)
		}
		header.DocFreq += 1
		header.Entries[docId] = &PostingEntry{
			DocId: docId,
			Location: loc,
		}
	}
	return nil
}

// if keyword not in index, return nil todo
func (h *InvertedIndex) GetEntriesFromKeyword(keyword string) []*PostingEntry{
	return nil
}

//todo
func (h *InvertedIndex) GetHeaderFromKeyword(keyword string) *Header{
	return nil
}



func (h *InvertedIndex) GetDocumentFreq(word string) (int, error) {
	if headerObject, err := h.HeadersTrie.GetObjectFromTrie(word); err == nil{
		return headerObject.(*Header).DocFreq, nil
	} else {
		return -1, errors.New("the word is not found.\n")
	}
}

func (pl *PostingList) String() string {
	ret := "|"
	for _, e := range pl.Entries{
		ret = ret + "D" + strconv.Itoa(e.DocId) + ":"
		str1 := make([]string, 0)
		for _, l := range e.Location{
			str1 = append(str1, strconv.Itoa(l))
		}
		ret = ret + strings.Join(str1, ",") + "|"
	}
	return ret
}