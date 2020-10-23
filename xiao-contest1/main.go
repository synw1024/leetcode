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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	solve(arr)
}

func solve(arr []int) {
	// counter := [][]int{}
	i, j, k, l, sum := 0, 1, 2, len(arr), 0
	for ; i < l-2; i++ {
		for j = i + 1; j < l-1; j++ {
			for k = j + 1; k < l; k++ {
				if arr[i] != arr[j] && arr[i] != arr[k] && arr[j] != arr[k] {
					if isTriangle(arr[i], arr[j], arr[k]) {
						sum++
						// counter = append(counter, []int{arr[i], arr[j], arr[k]})
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func isTriangle(i, j, k int) bool {
	if i+j > k && i+k > j && j+k > i {
		return true
	}
	return false
}

func has(counter [][]int, i, j, k int) bool {
	for _, v := range counter {
		if contains(v, i) && contains(v, j) && contains(v, k) {
			return true
		}
	}
	return false
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
