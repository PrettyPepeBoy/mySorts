package main

import (
	"dima/quicksort"
	"fmt"
)

func main() {
	a := []int{15, 7, 24, 56, 32, 1, 93, 43, 54, 22, 17}
	fmt.Println(a)
	quicksort.Sort(a, 0, len(a)-1)
	fmt.Println(a)

	b := []int{15, 7, 24, 56, 32, 1, 93, 43, 54, 22, 17}
	fmt.Println(b)
	quicksort.MySort(b, 0, len(b)-1)
	fmt.Println(b)
}
