package main

import (
	"fmt"
)

type Matrix [][]int

func main() {
	// Define two 2D arrays (matrices) with equal dimensions
	matrixA := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrixB := Matrix{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	// Perform matrix addition
	result := addMatrices(matrixA, matrixB)

	// Print the result matrix
	fmt.Println("Resultant Matrix after Addition:")
	fmt.Println(result)
}

func addMatrices(matrixA, matrixB Matrix) Matrix {
	nRows := len(matrixA)
	nCols := len(matrixA[0])

	// Create a result matrix to store the sum of matrixA and matrixB
	result := make(Matrix, nRows)
	for i := range result {
		result[i] = make([]int, nCols)
	}

	// Perform matrix addition using nested loops
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}

	return result
}

func (m Matrix) Format(f fmt.State, c rune) {
	for _, row := range m {
		fmt.Fprintf(f, "%v\n", row)
	}
}
