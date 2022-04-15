package main

import (
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
