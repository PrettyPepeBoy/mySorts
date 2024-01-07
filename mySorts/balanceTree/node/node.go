package node

type MySlc struct {
	Mas []int
}

func (a *MySlc) Init(val int) {
	a.Mas = append(a.Mas, val)
}

func (a *MySlc) Put(val int, i int) {
	root := a.Mas[i]

	if val < root {
		a.Mas[2*i+1] = val
	} else {
		a.Mas[2*i+2] = val
	}
}

func MyAppend(list []int, variable int) []int {
	var newList []int

	listLen := len(list) + 1

	if listLen <= cap(list) {
		newList = list[:listLen]
	} else {
		listCap := listLen
		if listCap < 2*len(list) {
			listCap = 2 * len(list)
		}

		newList = make([]int, listLen, listCap)
		copy(newList, list)
	}

	newList[len(list)] = variable

	return newList
}

//func ()

/*func (t Tree) Put(val int) {
	ptr := t.root
	j := 0
	for j < len(ptr.Mas) && val > ptr.Mas[j]{
		j++
	}
}*/
