package main

import (
	"bufio"
	"fmt"
	"github.com/stevenyao/go-opencc"
	"io"
	"os"
)

const (
	config_s2t = "/usr/share/opencc/s2t.json"
	config_t2s = "/usr/share/opencc/t2s.json"
)

func simple2tradition(input string) string {
	c := opencc.NewConverter(config_s2t)
	defer c.Close()
	return c.Convert(input)
}

func convertPipe(reader *bufio.Reader) {
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		text := simple2tradition(input)
		fmt.Printf("%s", text)
	}
}

func main() {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		var input string
		if os.Args[1] == "-h" {
			fmt.Println("Convert Simple Chinese font to Traditional font.")
		} else {
			for _, arg := range os.Args[1:] {
				input += arg
			}

			output := simple2tradition(input)
			fmt.Println(output)
		}

	} else {
		reader := bufio.NewReader(os.Stdin)
		convertPipe(reader)
	}
}
