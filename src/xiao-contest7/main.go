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

	temp := strings.Split(readLine(reader), " ")
	tempN, err := strconv.ParseInt(temp[0], 10, 64)
	checkError(err)
	n := int(tempN)
	tempV, err := strconv.ParseInt(temp[1], 10, 64)
	checkError(err)
	V := int(tempV)

	vw := [][]int{}
	for i := 0; i < n; i++ {
		temp := strings.Split(readLine(reader), " ")
		tempV, err := strconv.ParseInt(temp[0], 10, 64)
		checkError(err)
		tempW, err := strconv.ParseInt(temp[1], 10, 64)
		checkError(err)
		vw = append(vw, []int{int(tempV), int(tempW)})
	}

	res := solve(V, vw)
	fmt.Println(res)
}

func solve(V int, vw [][]int) int {
	var dp func(V int, vw [][]int) int
	dp = func(V int, vw [][]int) int {
		max := 0
		for i, v := range vw {
			if v[0] > V {
				continue
			}
			temp := append([][]int{}, vw[:i]...)
			_vw := append(temp, vw[i+1:]...)
			val := dp(V-v[0], _vw) + v[1]
			if val > max {
				max = val
			}
		}
		return max
	}
	return dp(V, vw)
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
