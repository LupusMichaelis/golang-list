package main

import (
	"fmt"
	"os"
)

func main() {
	var argList = List{nil}
	argList.AddMany(os.Args[1:]...)

	fmt.Println(argList)
	argList.Reverse()
	fmt.Println(argList)
}
