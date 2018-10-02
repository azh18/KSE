package KSE

import (
	"container/heap"
	"strings"
	"unicode"
)

const ValidLength = 4

type Document struct{
	DocId int
	Content string
	Words map[string]int
	wordWithMaxWeight *PriorityQueue
	Magnitude float64
	NUniqueKeyword int
	MaxTf int
}

// build new document and write words & docId & nUniqueWords with pre-processing
func NewDocument(content string, docId int) *Document{
	newDoc := &Document{
		DocId: docId,
		Content: content,
		Words: map[string]int{},
		wordWithMaxWeight: NewPriorityQueue(),
		Magnitude: 0.0,
		NUniqueKeyword: 0,
		MaxTf: 0,
	}
	words := map[string]int{}

	checkLetterFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	keywords := strings.FieldsFunc(content, checkLetterFunc) // trim all punctuations, spaces...
	for _, s := range keywords{
		s = strings.TrimRight(s, "s") // trim 's' at the tail of words
		s = strings.ToLower(s)
		if len(s) < ValidLength{ // omit words with length smaller than 4
			continue
		}
		// s in valid, add it to keywords
		if _, ok := words[s]; !ok{
			words[s] = 0
		} else {
			words[s] += 1
		}
	}

	for k := range words{
		if words[k] > newDoc.MaxTf{
			newDoc.MaxTf = words[k]
		}
		if words[k] == 1{
			newDoc.NUniqueKeyword += 1
		}
	}
	newDoc.Words = words
	return newDoc
}

// update Magnitude field and wordWithMaxWeight field when knowing document frequency of all words
func (d *Document) UpdateDocMetaWithDf(docFreq map[string]int){
	sumWeightSquare := 0.0
	for k, v := range d.Words{
		idf := 1.0/float64(docFreq[k])
		weight := float64(v)/float64(d.MaxTf) * idf
		sumWeightSquare += weight * weight
		// update wordWithMaxWeight, considering how to implement a priority queue with the fixed length
		newItem := NewPQItem(weight, k)
		heap.Push(d.wordWithMaxWeight, newItem)
	}
	d.Magnitude = sumWeightSquare / float64(len(d.Words))
}

func (d *Document) GetMagnitude() float64{
	return d.Magnitude
}

func (d *Document) GetTop5WeightWords() ([]string, error) {
	top5Items, err := d.wordWithMaxWeight.PopTopK(5)
	if err != nil{
		return nil, err
	}
	ret := make([]string, 0)
	for _, v := range top5Items{
		ret = append(ret, v.value.(string))
	}
	return ret, nil
}

func (d *Document) GetNUniqueWords() int{
	return d.NUniqueKeyword
}
