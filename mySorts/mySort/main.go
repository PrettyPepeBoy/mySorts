package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

type Edb []Employee

func (emp Edb) Swap(i, j int) Edb {
	emp[i], emp[j] = emp[j], emp[i]
	return emp
}

func (emp Edb) Check(i, j int) bool {
	if emp[i].Age < emp[j].Age {
		return emp[i].Age < emp[j].Age
	}
	return false
}

func (emp Edb) MySort() Edb {
	i := 0
	for i < len(emp) {
		for j := i; j < len(emp); j++ {
			if emp.Check(i, j) {
				emp.Swap(i, j)
			}
		}
		i++
	}
	return emp
}

func main() {
	emp := Edb{
		{"Dan", 15},
		{"Yan", 14},
		{"Kate", 54},
		{"Fenix", 34},
		{"Elizabeth", 22},
		{"Jack", 74},
		{"Jacob", 16},
	}
	fmt.Println(emp)

	fmt.Println(emp.MySort())
	fmt.Printf("%#v\n", emp)

}
