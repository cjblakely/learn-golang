package main

import "fmt"

func main() {
	xs := []float64{98,93,77,82,83}

	//total := 0.0
	//for _, v := range xs {
	//	total += v
	//}
	//fmt.Println(total / float64(len(xs)))
	fmt.Println(average(xs))


	//variadic function
	fmt.Println(add(1,2,3))

	//closure
	sub := func(x, y int) int {
		return x - y
	}
	fmt.Println(sub(1,1))

	xc := 0
	increment := func() int {
		xc++
		return xc
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	//recursion
	fmt.Println(factorial(4))

	//defer
	defer second()
	first()

}

func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	xf := uint(0)
	return func() (ret uint) {
		ret = xf
		xf += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}