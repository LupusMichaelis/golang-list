package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	var list = List{nil}
	for _, arg := range args {
		list.Add(arg)
	}

	fmt.Println(&list)
	list.Reverse()
	fmt.Println(&list)

	return
}
