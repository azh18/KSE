package query

import "github.com/zbw0046/KSE/dal"

type QueryResultItem struct{
	docId int
	Top5PostingLists map[string]*dal.PostingList
	NUniqueWordsInDoc int
	MagDoc float64
	SimilarityScore float64
}

type QueryResult dal.PriorityQueue

// todo
func ExecuteQuery(keyword string) QueryResult{
	return nil
}

// todo
func (q *QueryResult) String() string {
	return ""
}

// todo
func (qi *QueryResultItem) String() string{
	return ""
}