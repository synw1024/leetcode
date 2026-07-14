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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	cache := map[int]int{}
	var dp func(n int) int
	dp = func(n int) int {
		v, ok := cache[n]
		if ok {
			return v
		}

		if n <= 1 {
			return 1
		}
		min := math.MaxInt64
		for i := 1; true; i++ {
			m := int(math.Pow(float64(6), float64(i)))
			if m > n {
				break
			} else if m == n {
				return 1
			} else {
				res := dp(n-m) + 1
				if res < min {
					min = res
				}
			}
		}

		for i := 1; true; i++ {
			m := int(math.Pow(float64(9), float64(i)))
			if m == n {
				return 1
			} else if m > n {
				break
			}
			res := dp(n-m) + 1
			if res < min {
				min = res
			}
		}

		res := dp(n-1) + 1

		if res < min {
			min = res
		}
		cache[n] = min
		return min
	}
	return dp(n)
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
