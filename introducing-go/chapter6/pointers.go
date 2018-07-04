package main

import "fmt"

func main() {
	//original
	x := 5
	zero(&x)
	fmt.Println(x)

	//new func
	newPtr := new(int)
	one(newPtr)
	fmt.Println(*newPtr)
}

func zero(xPtr *int) {
	*xPtr = 0
}

func one(newPtr *int) {
	*newPtr = 1
}