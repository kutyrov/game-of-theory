package main

import (
	"errors"
)

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

// функия находит обратную матрицу
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
