package main

import "fmt"
//had to use go get github.com/jessevdk/go-flags
import flags "github.com/jessevdk/go-flags"

var opts struct {
	Name string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool `short:"s" long:"spanish" description:"Use Spanish Language"`
}

func main() {
	flags.Parse(&opts)

	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}