package main

import (
	"fmt"
	"strings"
)

func main() {
	//smaller string in bigger string
	fmt.Println(strings.Contains("test", "es"))

	//count number of times smaller string occurs
	fmt.Println(strings.Count("test", "t"))

	//if bigger starts with smaller
	fmt.Println(strings.HasPrefix("test", "te"))

	//bigger string ends with smaller
	fmt.Println(strings.HasSuffix("test", "st"))

	//find position of smaller string in bigger
	fmt.Println(strings.Index("test", "e"))

	//take a list of strings and join them together
	fmt.Println(strings.Join([]string{"a","b"}, "-"))

	//repeat a string
	fmt.Println(strings.Repeat("a", 5))

	//replace smaller string n bugger string with other string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))

	//split a string into a list of strings by a separating string
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	//convert to lowercase
	fmt.Println(strings.ToLower("TEST"))

	//covert to uppercase
	fmt.Println(strings.ToUpper("test"))

	//covert to a slice of bytes or vice versa
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t','e','s','t'})
	fmt.Println(str)
}
