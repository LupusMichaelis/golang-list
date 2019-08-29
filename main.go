package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	var argList = List{nil}
	for _, arg := range args {
		argList.Add(arg)
	}

	fmt.Println(&argList)
	argList.Reverse()
	fmt.Println(&argList)

	var numberList = List{nil}
	for _, arg := range []int{1, 2, 3, 4} {
		numberList.Add(arg)
	}

	fmt.Println(&numberList)
	numberList.Reverse()
	fmt.Println(&numberList)

	var anyList = List{nil}
	anyList.Add(1)
	anyList.Add("Yolo")
	anyList.Add(3.5)
	anyList.Add(0.5 + 1i + 3.5 - 2i)

	fmt.Println(&anyList)
	anyList.Reverse()
	fmt.Println(&anyList)
}
