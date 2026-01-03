package main

import "fmt"

func numeric(){
	var num32 int32 = 2000
	var num64 int64 = int64(num32)

	fmt.Println(num64)
}

func strings(){
	name := "RascaL"
	var c uint8 = name[0]
	var str = string(c)

	fmt.Println(c)
	fmt.Println(str)
}

func main(){
	strings()
}
