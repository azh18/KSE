package main

import (
	"bufio"
	"fmt"
	"github.com/zbw0046/KSE/KSE"
	"os"
)

func main(){
	filename := "collection-100.txt"
	db := KSE.NewDatabase(filename)
	queryFileName := "query-10.txt"
	queryKeyWords := ParseQueryFile(queryFileName)
	for _, k:=range queryKeyWords{
		fmt.Printf("query:%s\n", k)
		fmt.Printf("------------------------------------\n")
		ret := db.Query(k)
		fmt.Println(ret)
	}
	return
}

func ParseQueryFile(filename string) []string {
	ret := make([]string, 0)
	if fd, err := os.Open(filename); err != nil{
		panic("filename cannot be open.")
	} else {
		scanner := bufio.NewScanner(fd)
		for scanner.Scan(){
			keywords := scanner.Text()
			ret = append(ret, keywords)
		}
	}
	return ret
}