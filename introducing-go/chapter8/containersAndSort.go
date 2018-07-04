package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	//container/list - doubly linked list
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}

	//sort
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

type Person struct {
	Name string
	Age int
}

//create new type to sort by name
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

//create new type to sort by age
type ByAge []Person
func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}