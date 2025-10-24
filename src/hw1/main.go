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
	line := readLine(reader)
	temp := strings.Split(line, " ")
	var nums []int64
	for _, s := range temp {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	var min int64 = math.MaxInt64
	var min1 int64
	var min2 int64
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := int64(math.Abs(float64(nums[i] + nums[j])))
			if sum < min {
				min = sum
				min1 = nums[i]
				min2 = nums[j]
			}
		}
	}
	if min1 < min2 {
		fmt.Println(min1, min2, min)
	} else {
		fmt.Println(min2, min1, min)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
