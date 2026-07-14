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
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	nums := make([]string, 60)
	nums[0] = "1"
	for i := 1; i <= int(num); i++ {
		prev := nums[i-1]
		arr := []int{}
		for j := 0; j < len(prev); j++ {
			n, err := strconv.ParseInt(string(prev[j]), 10, 0)
			if err != nil {
				panic(err)
			}
			arr = add(arr, int(n))
		}
		for k := 0; k < len(arr); k += 2 {
			nums[i] += strconv.Itoa(arr[k+1]) + strconv.Itoa(arr[k])
		}
	}
	fmt.Println(nums[num])
}

func add(arr []int, n int) []int {
	l := len(arr)
	if l == 0 {
		arr = append(arr, n, 1)
		return arr
	}
	if arr[l-2] == n {
		arr[l-1]++
		return arr
	}
	arr = append(arr, n, 1)
	return arr
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
