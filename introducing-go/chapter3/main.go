package main

import "fmt"

var j string = "Hello, Scope"

func main() {
	//assign option 1
	var x string = "Hello, World"
	fmt.Println(x)

	//assign option 2
	var y string
	y = "Hello, World"
	fmt.Println(y)

	//value change examples
	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	var v string
	v = "first "
	fmt.Println(v)
	//v = v + "second"
	v += "second"
	fmt.Println(v)

	//equality
	var a string = "hello"
	var b string = "world"
	fmt.Println(a == b)

	var c string = "hello"
	var d string = "hello"
	fmt.Println(c == d)

	//shorthand
	e := 5
	fmt.Println(e)

	//scope
	fmt.Println(j)

	//const
	const u string = "Hello, Const"
	fmt.Println(u)

	//mass assignment
	var (
		f = 5
		g = 10
		h = 15
	)
	fmt.Println(f, g, h)
}

func f() {
	//compiles due to scope
	fmt.Println(j)
}
