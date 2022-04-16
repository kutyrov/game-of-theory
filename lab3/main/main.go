package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const crossroadPath = "data/crossroad.txt"
const disputePath = "data/dispute.txt"
const prisonersPath = "data/prisoners.txt"
const defaultPath = "data/12.txt"
const maxNumber = 50
const matrixSize = 10
const defaultMatrixSize = 2
const eps = 0.1

func crossing(a, b []pair) []pair {
	//to do
	return nil
}

func printCells(matrix [][]Win, data []pair) {
	for index := range data {
		fmt.Printf(
			"(%d %d)\n",
			int(matrix[data[index].first][data[index].second].AWin),
			int(matrix[data[index].first][data[index].second].BWin),
		)
	}
}

func getMatrix(path string) ([][]Win, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	matrix := make([][]Win, defaultMatrixSize)
	for jindex := range matrix {
		matrix[jindex] = make([]Win, defaultMatrixSize)
	}

	index := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())

		for jindex, word := range numbers {
			data := strings.Split(word, ",")
			matrix[index][jindex].AWin, _ = strconv.ParseFloat(data[0], 64)
			matrix[index][jindex].BWin, _ = strconv.ParseFloat(data[1], 64)
		}
		index += 1
	}

	return matrix, nil
}

func main() {
	fmt.Printf("Генерируем матрицу %d*%d\n\n", matrixSize, matrixSize)
	matrix := generateMatrix(matrixSize, matrixSize)
	printWins(matrix)
	resNash := eqNash(matrix)
	fmt.Printf("\nНаходим равновесие по Нэшу\n")
	if len(resNash) == 0 {
		fmt.Println("Решений нет")
	}
	printCells(matrix, resNash)
	fmt.Printf("\nНаходим оптимум по Парето\n")
	printCells(matrix, eqPareto(matrix))

	fmt.Println("\n\nПроверяем функции на игре \"Перекресток\"")
	matrix, err := getMatrix(crossroadPath)
	if err != nil {
		panic(err)
	}
	matrix[0][1].AWin -= eps
	matrix[1][0].BWin -= eps
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("\nНаходим оптимум по Парето\n")
	printCells(matrix, eqPareto(matrix))

	fmt.Println("\n\nПроверяем функции на игре \"Семейный спор\"")
	matrix, err = getMatrix(disputePath)
	if err != nil {
		panic(err)
	}
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("\nНаходим оптимум по Парето\n")
	printCells(matrix, eqPareto(matrix))

	fmt.Println("\n\nПроверяем функции на игре \"Дилемма заключённого\"")
	matrix, err = getMatrix(prisonersPath)
	if err != nil {
		panic(err)
	}
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("\nНаходим оптимум по Парето\n")
	printCells(matrix, eqPareto(matrix))

	fmt.Println("\n\nРешаем игру своего варианта")
	matrix, err = getMatrix(defaultPath)
	if err != nil {
		panic(err)
	}
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("\nНаходим оптимум по Парето\n")
	printCells(matrix, eqPareto(matrix))
}
