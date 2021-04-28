package main

import "fmt"

func main() {
	sudoku := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := isValidSudoku(sudoku)
	fmt.Println(res)
}

func isValidSudoku(board [][]byte) bool {
	var hashMap, hashMap2, hashMap3 [9]map[byte]bool
	for i := 0; i < 9; i++ {
		hashMap[i] = make(map[byte]bool)
		hashMap2[i] = make(map[byte]bool)
		hashMap3[i] = make(map[byte]bool)
	}

	for i, v := range board {
		for j, vv := range v {
			if vv == '.' {
				continue
			}

			if _, ok := hashMap[i][vv]; ok {
				return false
			} else {
				hashMap[i][vv] = true
			}

			if _, ok := hashMap2[j][vv]; ok {
				return false
			} else {
				hashMap2[j][vv] = true
			}

			if _, ok := hashMap3[i/3*3+j/3][vv]; ok {
				return false
			} else {
				hashMap3[i/3*3+j/3][vv] = true
			}
		}
	}
	return true
}
