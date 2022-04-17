package main

import (
	"errors"
	"math/rand"
	"time"
)

type Win struct {
	AWin float64
	BWin float64
}

// генерирует случайное целое число в интервале [-n,n)
func randInt(data int) int {
	return rand.Intn(2*data) - data
}

//Если в матрице одна строка преобразует её в слайс
func matrixToSlice(matrix [][]float64) []float64 {
	if len(matrix) == 1 {
		return matrix[0]
	} else {
		return nil
	}
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

func generateU(n int) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = 1
	}
	return data
}

func sumMatrix(matrix [][]float64) float64 {
	sum := 0.0
	for _, row := range matrix {
		for _, cell := range row {
			sum += cell
		}
	}
	return sum
}

func matrixOnFloat(matrix [][]float64, value float64) [][]float64 {
	data := make([][]float64, len(matrix))
	for i := range data {
		data[i] = make([]float64, len(matrix[0]))
		for j := range matrix[0] {
			data[i][j] = value * matrix[i][j]
		}
	}
	return data
}

func winToMatrix(data [][]Win) ([][]float64, [][]float64) {
	matrixA := make([][]float64, len(data))
	matrixB := make([][]float64, len(data))
	for i := range matrixA {
		matrixA[i] = make([]float64, len(data))
		matrixB[i] = make([]float64, len(data))
	}

	for i, row := range data {
		for j, cell := range row {
			matrixA[i][j] = cell.AWin
			matrixB[i][j] = cell.BWin
		}
	}
	return matrixA, matrixB
}

func generateMatrix(row, col int) [][]Win {
	rand.Seed(time.Now().UnixNano())
	if row <= 0 || col <= 0 {
		return nil
	}
	matrix := make([][]Win, row)
	for index := range matrix {
		matrix[index] = make([]Win, col)
		for jindex := range matrix[index] {
			matrix[index][jindex].AWin = float64(randInt(maxNumber))
			matrix[index][jindex].BWin = float64(randInt(maxNumber))
		}
	}
	return matrix
}

// функция транспонирует матрицу
func TransposeMatrix(matrix [][]float64) [][]float64 {
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

// функция удаляет выбранные строку и стобец
func ReduceMatrix(matrix [][]float64, row, column int) ([][]float64, error) {
	// проверяем крайние случаи
	if row < 0 {
		return nil, errors.New("Номер строки должен быть больше нуля")
	}
	if column < 0 {
		return nil, errors.New("Номер столбца должен быть больше нуля")
	}
	if row > len(matrix) {
		return nil, errors.New("Слишком большой номер строки")
	}
	if column > len(matrix[0]) {
		return nil, errors.New("Слишком большой номер солбца")
	}

	// вводим новые переменные
	rowNew := len(matrix) - 1
	columnNew := len(matrix[0]) - 1
	matrixNew := make([][]float64, rowNew, rowNew)
	for i := range matrixNew {
		matrixNew[i] = make([]float64, columnNew, columnNew)
	}

	// проходимся по старой матрице, записывая нужные элементы в новую
	indexNew, jindexNew := 0, 0
	for index := range matrix {
		for jindex := range matrix[index] {
			if index != row && jindex != column {
				matrixNew[indexNew][jindexNew] = matrix[index][jindex]
				jindexNew += 1
				if jindexNew == columnNew {
					jindexNew = 0
					indexNew += 1
				}

			}
		}
	}

	return matrixNew, nil
}

// функция вычисляет определитель матрицы, разложением по первой строке
func DetMatrix(matrix [][]float64) (float64, error) {
	if len(matrix) != len(matrix[0]) {
		return 0, errors.New("матрица необратима")
	}
	if len(matrix) == 1 {
		return matrix[0][0], nil
	}
	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0], nil
	} else {
		det := 0.0
		for index := range matrix {
			temp, _ := ReduceMatrix(matrix, 0, index)
			detInter, _ := DetMatrix(temp)
			if index%2 == 0 {
				det += matrix[0][index] * detInter
			} else {
				det -= matrix[0][index] * detInter
			}

		}
		return det, nil
	}
}

// функция находит обратную матрицу
func InverseMatrix(matrix [][]float64) ([][]float64, error) {
	det, err := DetMatrix(matrix)
	if err != nil {
		return nil, err
	}

	transposeMatrix := TransposeMatrix(matrix)

	inverseMatrix := make([][]float64, len(transposeMatrix), len(transposeMatrix))
	for index := range inverseMatrix {
		inverseMatrix[index] = make([]float64, len(transposeMatrix[0]), len(transposeMatrix[0]))
	}

	for index := range inverseMatrix {
		for jindex := range inverseMatrix[index] {
			matrixInter, _ := ReduceMatrix(transposeMatrix, index, jindex)
			inverseMatrix[index][jindex], _ = DetMatrix(matrixInter)
			if (index+jindex)%2 == 1 {
				inverseMatrix[index][jindex] *= -1
			}
			inverseMatrix[index][jindex] /= float64(det)
		}
	}

	return inverseMatrix, nil

}
