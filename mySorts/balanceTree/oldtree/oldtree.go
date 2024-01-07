package oldtree

import (
	"fmt"
)

type Tree struct {
	Val     int
	left    *Tree
	right   *Tree
	parent  *Tree
	BFactor int
}

func CreateTree(value int) Tree {
	var t Tree
	t.Val = value
	t.BFactor = 0
	return t
}

func (t *Tree) rightTurn() *Tree {
	t.left.parent = t.parent
	t.parent.left = t.left
	t.left.BFactor--
	t.parent = t.left
	t.left = nil
	t.BFactor = 0
	fmt.Println(t.parent)
	fmt.Println(t.parent.left)
	return t
}

func (t *Tree) leftTurn() *Tree {
	t.right.parent = t.parent
	t.parent.right = t.right
	t.right.BFactor++
	t.parent = t.right
	t.right = nil
	t.BFactor = 0
	return t
}

func (t *Tree) setHeight() (bool, *Tree) {
	if t.parent == nil {
		if t.BFactor == 2 || t.BFactor == -2 {
			return true, t
		} else {
			return false, t
		}
	}

	if t.Val > t.parent.Val {
		t.parent.BFactor--
		if t.parent.BFactor == -2 {
			return true, t.parent
		}
		if t.parent.BFactor == 0 {
			return false, t
		} else {
			return t.parent.setHeight()
		}
	} else {
		t.parent.BFactor++
		if t.parent.BFactor == 2 {
			return true, t.parent
		}
		if t.parent.BFactor == 0 {
			return false, t
		} else {
			return t.parent.setHeight()
		}
	}
}

func (t *Tree) Put(value int) bool {
	var check bool
	if t.Val > value {
		if t.left == nil {
			t.left = &Tree{
				Val:     value,
				parent:  t,
				BFactor: 0,
			}
			check, t = t.left.setHeight()
			if check {
				if t.BFactor == 2 {
					t.rightTurn()
				} else {
					fmt.Println("неудача")
				}
			}
			return false
		}
		return t.left.Put(value)
	}

	if t.Val < value {
		if t.right == nil {
			t.right = &Tree{
				Val:     value,
				parent:  t,
				BFactor: 0,
			}
			check, t = t.right.setHeight()
			if check {
				fmt.Println(t)
			}
			return false
		}
		return t.right.Put(value)
	}

	return true
}

func (t *Tree) FindMax() *Tree {
	if t.right == nil {
		return t
	}
	return t.right.FindMax()
}

func (t *Tree) FindMin() *Tree {
	if t.left == nil {
		return t
	}
	return t.left.FindMin()
}

func (t *Tree) FindElem(value int) *Tree {

	if t.Val > value {
		if t.left != nil {
			return t.left.FindElem(value)
		} else {
			return nil
		}
	}

	if t.Val < value {
		if t.right != nil {
			return t.right.FindElem(value)
		} else {
			return nil
		}
	}

	return t
}

func (t *Tree) Delete(value int) bool {
	node := t.FindElem(value)
	if node == nil {
		return false
	}

	if node.right == nil {
		if node.parent.Val > node.Val {
			node.parent.left = node.left
			if node.left != nil {
				node.left.parent = node.parent
			}
		} else {
			node.parent.right = node.left
			if node.left != nil {
				node.left.parent = node.parent
			}
		}
		return true
	}

	minNode := node.right.FindMin()
	if minNode.parent.Val != node.Val {
		node.Val = minNode.Val
		minNode.parent.left = minNode.right
		if minNode.right != nil {
			minNode.right.parent = minNode.right
		}
		return true
	}
	node.Val = minNode.Val
	node.right = minNode.right
	if minNode.right != nil {
		minNode.right.parent = node
	}
	return true
}

func (t *Tree) Print() {
	if t == nil {
		return
	}
	t.left.Print()
	fmt.Println("this node")
	fmt.Println(t.Val)
	fmt.Println(t.BFactor)
	t.right.Print()
}
