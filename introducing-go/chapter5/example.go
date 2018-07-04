package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	//for i := 0; i < 5; i++ {
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	//fmt.Println(total / 5)
	fmt.Println(total / float64(len(x)))

	//for loop with current position
	var total2 float64 = 0

	//throws error due to i not used
	//for i, value := range x {
	//	total2 += value
	//}

	for _, value := range x {
		total2 += value
	}

	//map
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"

	fmt.Println(elements["Li"])

	//does not exist, returns zero value of type
	fmt.Println(elements["Un"])

	//better way to check
	name, ok := elements["Un"]
	fmt.Println(name, ok) //ok set to false

	//often seen as
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	//shorthand
	elements2 := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
	}
	fmt.Println(elements2["H"])

	//map in map
	elements3 := map[string]map[string]string{
		"H" : map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He" : map[string]string{
			"name":"Helium",
			"state":"gas",
		},
	}

	if el, ok := elements3["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
