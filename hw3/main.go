package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	s := readLine(reader)
	temp, err := strconv.ParseInt(s, 10, 0)
	n := int(temp)
	if err != nil {
		panic(err)
	}
	cmds := []string{}
	for i := 0; i < n; i++ {
		s := readLine(reader)
		temp := strings.Split(s, "=")
		cmds = append(cmds, temp...)
	}

	cache := make([]int, 100)
	for i := 0; i < len(cmds); i += 2 {
		cmd := cmds[i]
		temp, err := strconv.ParseInt(cmds[i+1], 10, 0)
		size := int(temp)
		if err != nil {
			panic(err)
		}
		if size == 0 && cmd == "REQUEST" {
			fmt.Println("error")
		}
		if cmd == "REQUEST" {
			count := 0
			j := 0
			for j < len(cache) {
				if cache[j] == 0 {
					count++
					if count == size {
						index := j + 1 - size
						cache[index] = size
						fmt.Println(index)
						break
					}
					j++
					continue
				}
				count = 0
				j += cache[j]
			}
			if j >= len(cache) {
				fmt.Println("error")
			}
		} else {
			if size < 0 || size >= 100 {
				fmt.Println("error")
			} else if cache[size] == 0 {
				fmt.Println("error")
			} else {
				cache[size] = 0
			}
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
