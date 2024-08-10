package main

func main() {
	println(Determinant([][]int{{6, 1, 1}, {4, -2, 5}, {2, 8, 7}}))
}

func Determinant(matrix [][]int) int {
	//base case for a 1x1 matrix
	if len(matrix) == 1 {
		return matrix[0][0]
	}
	//base case for a 2x2 matrix
	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}
	//case of a 3x3 matrix or larger
	det := 0
	for col := 0; col < len(matrix); col++ {
		// Create the submatrix by removing the first row and the current column
		submatrix := make([][]int, len(matrix)-1)
		for i := range submatrix {
			submatrix[i] = make([]int, len(matrix)-1)
		}
		//populating the sub-matrix
		for i := 1; i < len(matrix); i++ {
			k := 0
			for j := 0; j < len(matrix); j++ {
				if j != col {
					submatrix[i-1][k] = matrix[i][j]
					k++
				}
			}
		}
		// Alternate sign for cofactor expansion
		sign := 1
		if col%2 != 0 {
			sign = -1
		}
		det += sign * matrix[0][col] * Determinant(submatrix)
	}
	return det
}
