package main

import "fmt"

const maxRand = 1000
const maxSlice = 20
const matrixSize = 10
const maxEpsilon = 0.000001
const numberPlayers = 2
const firstPlayerWin = 101
const secondPlayerWin = -101

func main() {
	matrix := generateMatrix(matrixSize)
	printMatrix(matrix)

	x := randomSlice(matrixSize, maxSlice)

	fmt.Println()
	fmt.Println("Вектор мнений:")
	printSlice(x)

	steps, x := stepFirst(matrix, x)

	fmt.Println()
	fmt.Println("Итоговый вектор мнений")
	printSlice(x)
	fmt.Printf("Решение найдено за %d шагов\n\n", steps)

	stepSecond(matrix, numberPlayers)
}
