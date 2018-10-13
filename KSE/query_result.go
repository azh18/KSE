package KSE

import (
	"container/heap"
	"errors"
	"github.com/zbw0046/KSE/dal"
	"github.com/zbw0046/KSE/tools"
	"sort"
	"strconv"
)

type Query struct{
	keywords string
	db *Database
	rankResult *QueryResult
	kValue int
}

type QueryResultItem struct{
	docId int
	Top5PostingLists map[string]*dal.PostingList
	NUniqueWordsInDoc int
	MagDoc float64
	SimilarityScore float64
}

type QueryResult = dal.PriorityQueue

// for each keyword with the product of weights in the query and the document
type wordPartialProduct struct{
	keyword string
	partialProduct float64
}

func NewQuery(keywords string, db *Database, kValue int) *Query{
	if db == nil{
		return nil
	}
	return &Query{
		keywords:keywords,
		db:db,
		kValue:kValue,
	}
}

// find query result and sort by similarity score
func (q *Query) Execute() error {
	ret := dal.NewPriorityQueue(10000, false)
	keywords := tools.TextPreProcess(q.keywords)
	// 0. compute vector of query by tf*idf, where idf is based on the whole db
	queryVector := GetQueryVector(keywords, q.db)
	db := q.db
	// 1. find docs that have same keywords with the help of inverted index
	// docs do not have same keywords will have a zero inner product, meaning that will not in the ranking
	// for each match doc, store its id and corresponding matching words with partial product for computing inner product
	matchDoc := map[int][]wordPartialProduct{}
	for k:= range queryVector {
		docsEntrys := db.index.GetEntriesFromKeyword(k)
		if docsEntrys == nil{
			// keyword not in index, do not compute anymore
			continue
		}
		for _, docEntry := range docsEntrys{
			var df int
			if df_, err := db.index.GetDocumentFreq(k); err != nil{
				// should able to find k, because if k not in index, it should not have docEntrys
				return errors.New("Query Failed. no this word in index.\n")
			} else {
				df = df_
			}
			wordWeightInDoc := tools.ComputeNormalizeTFIDF(len(docEntry.Location), db.docs[docEntry.DocId].MaxTf, df, len(db.docs))
			product := queryVector[k] * wordWeightInDoc
			matchDoc[docEntry.DocId] = append(matchDoc[docEntry.DocId], wordPartialProduct{
				keyword: k,
				partialProduct: product,
			})
		}
	}
	// 2. for each docs, compute cosine similarity and put it into priority queue
	for docId := range matchDoc{
		innerProduct := 0.0
		for idx := range matchDoc[docId]{
			innerProduct += matchDoc[docId][idx].partialProduct
		}
		division := db.docs[docId].GetMagnitude() * queryVector.GetMagnitude()
		cosineSimilarity := innerProduct / division
		top5WordsInDoc, err := db.docs[docId].GetTop5WeightWords()
		if err != nil{
			return errors.New("Query Failed. get top-5 words error.\n")
		}
		top5List := make(map[string]*dal.PostingList, 0)
		for _, w := range top5WordsInDoc{
			top5List[w] = db.index.GetHeaderFromKeyword(w).PostingList
		}
		queryResultItem := &QueryResultItem{
			docId: docId,
			Top5PostingLists: top5List,
			NUniqueWordsInDoc: db.docs[docId].GetNUniqueWords(),
			MagDoc: db.docs[docId].GetMagnitude(),
			SimilarityScore: cosineSimilarity,
		}
		pqItem := dal.NewPQItem(cosineSimilarity*100000+float64(docId)/100000, queryResultItem)
		heap.Push(ret, pqItem)
	}
	q.rankResult = ret
	return nil
}

// get query vector based on the v-space of documents
func GetQueryVector(keywords []string, db *Database) dal.DocVector{
	qVector := make(dal.DocVector, 0)
	maxTf := 0
	tfs := make(map[string]int, 0)
	for _, k := range keywords{
		if _, ok := tfs[k]; !ok{
			tfs[k] = 0
		}
		tfs[k]++
		if tfs[k] > maxTf {
			maxTf = tfs[k]
		}
	}
	for k := range tfs{
		df, err := db.index.GetDocumentFreq(k)
		// omit the keywords that have zero document frequency
		if err != nil{
			continue
		}
		qVector[k] = tools.ComputeNormalizeTFIDF(tfs[k], maxTf, df, len(db.docs))
	}
	return qVector
}

// get top-5 results and print it to return
func (q *Query) GetResult() string {
	ret := ""
	allItems, err:= q.rankResult.PopTopK(q.kValue)
	if err != nil{
		panic("panic")
	}
	for _, item := range allItems{
		qResItem := item.GetValue().(*QueryResultItem)
		ret += qResItem.String()
		ret += "\n"
	}
	return ret
}

// print one result in the format given in homework
func (qi *QueryResultItem) String() string{
	ret := ""
	ret = ret + "DID: " + strconv.Itoa(qi.docId) + "\n"
	s := make([]string, 0)
	for k := range qi.Top5PostingLists{
		s = append(s, k)
	}
	sort.Strings(s)
	for _, k := range s{
		ret = ret + k + " -> "
		ret = ret + qi.Top5PostingLists[k].String() + "\n"
	}
	ret = ret + "Number of unique keywords in document: " + strconv.Itoa(qi.NUniqueWordsInDoc) + "\n"
	ret = ret + "Magnitude of the document vector (L2 norm): " + strconv.FormatFloat(qi.MagDoc, 'f', 2 ,64) + "\n"
	ret = ret + "Similarity score: " + strconv.FormatFloat(qi.SimilarityScore, 'f', 2 ,64) + "\n"
	return ret
}