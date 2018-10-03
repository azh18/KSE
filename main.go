package main

import "github.com/zbw0046/KSE/KSE"

func main(){
	filename := "a.txt"
	db := KSE.NewDatabase(filename)
	ret := db.Query("kkk")
	println(ret)
	return
}
