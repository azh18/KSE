package main

import (
	"bufio"
	"fmt"
	"github.com/zbw0046/KSE/KSE"
	"os"
	"strconv"
)

const(
	helpString = "Error: Parameters not valid!\n" +
		"Usage:\n./KSE.o <PATH TO DATA FILE> <PATH TO QUERY FILE> <K>\n" +
		"Example:\n./KSE.o collection-100.txt query-10.txt 3\n"
)

func main(){
	argNum := len(os.Args)
	if argNum != 4 {
		fmt.Printf(helpString)
		return
	}
	filename := os.Args[1]
	if fd, err := os.Open(filename); err != nil{
		fmt.Printf("Error: Data file %s does not exist!\n", filename)
		return
	} else {
		fd.Close()
	}
	db := KSE.NewDatabase(filename)
	queryFileName := os.Args[2]
	if fd, err := os.Open(queryFileName); err != nil{
		fmt.Printf("Error: Query file %s does not exist!\n", queryFileName)
		return
	} else {
		fd.Close()
	}
	queryKeyWords := ParseQueryFile(queryFileName)
	kValue, err := strconv.Atoi(os.Args[3])
	if err != nil || kValue <= 0{
		fmt.Printf("Error: k value %s is not valid!\n", os.Args[3])
		return
	}
	for _, k:=range queryKeyWords{
		fmt.Printf("query:%s\n", k)
		fmt.Printf("------------------------------------\n")
		ret := db.Query(k, kValue)
		fmt.Println(ret)
	}
	return
}

// Get the query keywords from the query file
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