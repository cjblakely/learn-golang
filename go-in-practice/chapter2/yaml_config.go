package main

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}