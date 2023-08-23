package main

import "fmt"

func main() {
	size := 3
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	matrix2 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	var sum [][]int
	sum = sumMatrix(matrix1, matrix2, size)
	fmt.Println("Sum Matrix:")
	printMatrix(sum)
	product := multiplyMatrix(matrix1, matrix2, size)
	fmt.Println("Product Matrix:")
	printMatrix(product)
}

func sumMatrix(matrix1 [][]int, matrix2 [][]int, size int) [][]int {
	var sumMatrix [][]int

	// Initialize the sumMatrix with appropriate dimensions
	for i := 0; i < size; i++ {
		row := make([]int, size)
		sumMatrix = append(sumMatrix, row)
	}

	// Perform matrix addition
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			sumMatrix[row][col] = matrix1[row][col] + matrix2[row][col]
		}
	}
	return sumMatrix
}
func multiplyMatrix(matrix1 [][]int, matrix2 [][]int, size int) [][]int {
	product := make([][]int, size)
	for i := 0; i < size; i++ {
		product[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			product[i][j] = 0
			for k := 0; k < size; k++ {
				product[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return product
}
func printMatrix(matrix [][]int) {
	fmt.Println("{")
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	fmt.Println("}")
}
