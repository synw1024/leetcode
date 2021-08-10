package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 512)
	s := readLine(reader)
	chart, _, err := reader.ReadLine()
	if err == io.EOF {
		panic(err)
	}
	c := rune(chart[0])

	count := 0
	for _, v := range s {
		if v == c || c == v-32 || c == v+32 {
			count++
		}
	}

	fmt.Println(count)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
