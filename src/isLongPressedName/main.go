package main

import "fmt"

func main() {
	res := isLongPressedName("laiden", "laiden")
	fmt.Println(res)
}

func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] {
		return false
	}
	i := 0
	for j := 0; j < len(typed); j++ {
		if i < len(name) && typed[j] == name[i] {
			i++
		} else if typed[j] != typed[j-1] {
			return false
		}
	}
	return i == len(name)
}
