package main

import "fmt"

func main() {
	a := new(Android)
	//name never actually set
	a.Person.Talk()
	a.Talk()
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	//Person Person //android "has" a person
	Person //android "is" a person
	Model string
}