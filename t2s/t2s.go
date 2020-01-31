package main

import (
	"fmt"
	"github.com/stevenyao/go-opencc"
	"os"
)

const (
	config_s2t = "/usr/share/opencc/s2t.json"
	config_t2s = "/usr/share/opencc/t2s.json"
)

func tradition2simple(input string) string {
	c := opencc.NewConverter(config_t2s)
	defer c.Close()
	return c.Convert(input)
}

func main() {

	var input string
	for _, arg := range os.Args[1:] {
		input += arg
	}

	if os.Args[1] == "-h" {
		fmt.Println("Convert Traditional Chinese font to Simple font.")
	} else {

		for _, arg := range os.Args[1:] {
			input += arg
		}

		output := tradition2simple(input)
		fmt.Println(output)
	}
}
