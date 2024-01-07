package main

import (
	"dima/tree"
)

func main() {
	t := tree.CreateTree(8)
	t.Put(10)
	t.Put(6)
	t.Put(4)
	t.Print()
	t.Delete(10)
	t.Print()

}
