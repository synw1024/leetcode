package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < 3; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := solve(arr)
	fmt.Println(int(math.Abs(float64(res))))
}

func solve(arr []int) int {
	pos, steps, size := arr[0], arr[1], arr[2]
	if pos >= 0 {
		for ; steps > 0; steps-- {
			if pos > 0 {
				pos -= size
			} else {
				pos += size
			}

			if pos < 0 {
				steps--
				break
			}
		}
	} else {
		for ; steps > 0; steps-- {
			if pos > 0 {
				pos -= size
			} else {
				pos += size
			}

			if pos > 0 {
				steps--
				break
			}
		}
	}

	if steps%2 == 0 {
		return pos
	}
	return size - int(math.Abs(float64(pos)))
}

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
