package KSE

import (
	"bufio"
	"fmt"
	"github.com/zbw0046/KSE/dal"
	"os"
)

type Database struct{
	docs map[int]*dal.Document // map from docId to the document
	index *dal.InvertedIndex // the inverted index with posting lists
}

// build db (documents + index) from document file
func NewDatabase(filename string) *Database{
	db := &Database{
		docs: make(map[int]*dal.Document),
		index: dal.NewInvertedIndex(),
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
			d.docs[docId] = dal.NewDocument(newLine, docId)
			docId++
		}
		if err := sc.Err(); err != nil{
			panic("read file content error.\n")
			return
		}
	}
	// phase 2: build inverted index and update metadata of documents
	for docId := range d.docs{
		d.index.AddWords(docId, d.docs[docId].GetWordsWithLocations())
	}
	for docId := range d.docs{
		if err := d.docs[docId].UpdateDocMetaWithDf(d.index.GetDocumentFreq, len(d.docs)); err != nil{
			panic("cannot find document frequency in inverted index. stop.\n")
		}
	}
	return
}

func (d *Database) Query(keywords string) string {
	q := NewQuery(keywords, d)
	if err := q.Execute(); err != nil{
		ret := fmt.Sprintf("Error occured while querying. err=\n%+=v\n", err)
		return ret
	}
	return q.GetResult()
}