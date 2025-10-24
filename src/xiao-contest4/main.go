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

	for i := 0; i < 2; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	solve(arr)
}

func solve(arr []int) {
	N, K, count := arr[0], arr[1], 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= N; k++ {
				if (i+j)%K == 0 && (i+k)%K == 0 && (j+k)%K == 0 {
					count++
				}
			}
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
