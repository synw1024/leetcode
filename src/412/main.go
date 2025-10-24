package main

import (
	"strconv"
)

func main() {

}

func fizzBuzz(n int) []string {
	temp := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			temp = append(temp, "FizzBuzz")
		} else if i%3 == 0 {
			temp = append(temp, "Fizz")
		} else if i%5 == 0 {
			temp = append(temp, "Buzz")
		} else {
			temp = append(temp, strconv.Itoa(i))
		}
	}
	return temp
}
