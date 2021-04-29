package main

func main() {
	matrix := [][]int{}
	rotate(matrix)
}

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		matrix[i], matrix[l-1-i] = matrix[l-1-i], matrix[i]
	}
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
