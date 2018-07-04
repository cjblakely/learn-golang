package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//file, err := os.Open("test.txt")
	//	//if err !=  nil {
	//	//	//handle the error here
	//	//	return
	//	//}
	//	//defer file.Close()
	//	//
	//	////get the file size
	//	//stat, err := file.Stat()
	//	//if err != nil {
	//	//	return
	//	//}
	//	////read the file
	//	//bs := make([]byte, stat.Size())
	//	//_, err = file.Read(bs)
	//	//if err != nil {
	//	//	return
	//	//}

	//shorter way to read a file
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	//creating a file
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("test")

	// get the contents of a directory
	//dir, err := os.Open(".")
	//if err != nil {
	//	return
	//}
	//defer dir.Close()
	//
	//fileInfos, err := dir.Readdir(-1)
	//if err != nil {
	//	return
	//}
	//
	//for _, fi := range fileInfos {
	//	fmt.Println(fi.Name())
	//}

	// better way to get everything recursively
	filepath.Walk(".", func(path string, info os.Fileinfo, err error) error {
		fmt.Println(path)
		return nil
	})
}