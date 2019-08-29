package main

import (
	"fmt"
	"lupusmic.org/golang-list/list"
	"os"
)

func main() {
	var argList = list.List{}
	argList.AddMany(os.Args[1:]...)

	fmt.Println(argList)
	argList.Reverse()
	fmt.Println(argList)
}
