package main

import (
	"math/rand"
	"time"
)

func printMatrix(matrix [][]float64) {
	if len(matrix) == 0 {
		return
	}
	for row := range matrix {
		printSlice(matrix[row])
	}
}

// генерирует стохастическую матрицу размера size*size
func generateMatrix(size int) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = normSlice(randomSlice(size, maxRand))
	}
	return matrix
}

func multSlice(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}
	data := 0.0
	for i := range a {
		data += a[i] * b[i]
	}
	return data
}

func getCol(matrix [][]float64, n int) []float64 {
	data := make([]float64, len(matrix))
	for i := range matrix {
		data[i] = matrix[i][n]
	}
	return data
}

func multMatrix(A, B [][]float64) [][]float64 {
	if len(A[0]) != len(B) {
		return nil
	}

	data := make([][]float64, len(A))
	for i := range data {

		data[i] = make([]float64, len(B[i]))
		for j := range B[i] {
			data[i][j] = multSlice(A[i], getCol(B, j))
		}
	}
	return data
}

func matrixToSlice(matrix [][]float64) []float64 {
	if len(matrix) == 1 {
		return matrix[0]
	} else {
		return nil
	}
}

func sliceToMatrix(data []float64) [][]float64 {
	matrix := make([][]float64, 1)
	matrix[0] = data
	return matrix
}

// функция транспонирует матрицу
func transposeMatrix(matrix [][]float64) [][]float64 {
	matrixNew := make([][]float64, len(matrix[0]), len(matrix[0]))
	for index := range matrixNew {
		matrixNew[index] = make([]float64, len(matrix), len(matrix))
	}
	for index := range matrix {
		for jindex := range matrix[0] {
			matrixNew[jindex][index] = matrix[index][jindex]
		}
	}
	return matrixNew
}
