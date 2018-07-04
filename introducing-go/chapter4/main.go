package main

import "fmt"

func main() {
	//for loop (only loop in go)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {

		//fmt.Println(i)

		//basic if statement
		if i % 2 == 0 {
			//even
			fmt.Println(i, "even")
		} else {
			//odd
			fmt.Println(i, "odd")
		}

		// else if
		if i % 2 == 0 {
			// divisible by 2
		} else if i % 3 == 0 {
			// divisible by 3
		} else if i % 4 == 0 {
			// divisible by 4
		}

		//switch
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
		}

	}

}
