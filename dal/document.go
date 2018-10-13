package dal

import (
	"container/heap"
	"github.com/zbw0046/KSE/tools"
	"math"
)

type WordLoc map[string][]int

type Document struct{
	DocId int
	Content string
	Words WordLoc
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
		Words: WordLoc{},
		wordWithMaxWeight: NewPriorityQueue(5, true),
		Magnitude: 0.0,
		NUniqueKeyword: 0,
		MaxTf: 0,
	}
	words := WordLoc{}
	keywords := tools.TextPreProcess(content)
	pos := 0
	for _, s := range keywords{
		if _, ok := words[s]; !ok{
			words[s] = make([]int, 0)
		}
		words[s]  = append(words[s], pos)
		pos++
	}

	for k := range words{
		if len(words[k]) > newDoc.MaxTf{
			newDoc.MaxTf = len(words[k])
		}
		if len(words[k]) == 1{
			newDoc.NUniqueKeyword += 1
		}
	}
	newDoc.Words = words
	return newDoc
}

// update Magnitude field and wordWithMaxWeight field when knowing document frequency of all words
func (d *Document) UpdateDocMetaWithDf(getDocFreq func (string) (int, error), nTotalDoc int) error {
	sumWeightSquare := 0.0
	for k, v := range d.Words{
		docFreq, err := getDocFreq(k)
		if err != nil{
			return err
		}
		weight := tools.ComputeNormalizeTFIDF(len(v), d.MaxTf, docFreq, nTotalDoc)
		sumWeightSquare += weight * weight
		// update wordWithMaxWeight, considering how to implement a priority queue with the fixed length
		priority := weight * 1000000 + GetStringPriority(k)/1000
		newItem := NewPQItem(priority, k)
		heap.Push(d.wordWithMaxWeight, newItem)
	}
	d.Magnitude = math.Sqrt(sumWeightSquare)
	return nil
}

func (d *Document) GetMagnitude() float64{
	return d.Magnitude
}

// Get top-5 words with the maximum weight
func (d *Document) GetTop5WeightWords() ([]string, error) {
	newQueue := d.wordWithMaxWeight.Copy()
	top5Items, err := newQueue.PopTopK(5)
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

func (d *Document) GetWordsWithLocations() (WordLoc){
	return d.Words
}

func GetStringPriority(s string) float64{
	return float64(s[3] + s[2]*26 + s[1]*26*26+ s[0]*26*26*26)/(26*26*26*26)
}