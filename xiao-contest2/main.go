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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < 3; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := solve(arr)
	if res < 0 {
		fmt.Println(-res)
	} else {
		fmt.Println(res)
	}
}

func solve(arr []int) int {
	pos, steps, size := arr[0], arr[1], arr[2]
	if pos < 0 {
		pos = -pos
	}

	if steps*size <= pos {
		return pos - steps*size
	}
	remain := steps*size - pos

	if (remain/size)%2 == 0 {
		return remain % size
	} else {
		return size - remain%size
	}
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
