package KSE

import (
	"bufio"
	"os"
)

type Database struct{
	docs map[int]*Document // map from docId to the document
	index *InvertedIndex // the inverted index with posting lists
}

// build db (documents + index) from document file
func NewDatabase(filename string) *Database{
	db := &Database{
		docs: make(map[int]*Document),
		index: NewInvertedIndex(),
	}
	db.BuildDBFromFile(filename)
	return db
}

func (d *Database) BuildDBFromFile(filename string){
	// phase 1: read data into documents database
	docId := 0
	if fd, err := os.Open(filename); err != nil{
		panic("file not found.\n")
		return
	} else {
		sc := bufio.NewScanner(fd)
		for sc.Scan(){
			newLine := sc.Text()
			if len(newLine) <= 1{
				continue
			}
			d.docs[docId] = NewDocument(newLine, docId)
			docId++
		}
		if err := sc.Err(); err != nil{
			panic("read file content error.\n")
			return
		}
	}
	// phase 2: build inverted index and update metadata of documents

}