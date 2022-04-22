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
	if len(a) == 0 || len(b) == 0 {
		return make([]pair, 0)
	}

	data := make([]pair, 0)
	for _, valueA := range a {
		for _, valueB := range b {
			if valueA == valueB {
				data = append(data, valueA)
			}
		}
	}
	return data
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
	fmt.Printf("Находим оптимум по Парето\n")
	reshPareto := eqPareto(matrix)
	printCells(matrix, reshPareto)
	cross := crossing(resNash, reshPareto)
	if len(cross) == 0 {
		fmt.Printf("\nПересечений решений нет\n")
	} else {
		fmt.Printf("\nПересечение решений\n")
		printCells(matrix, cross)
	}
	fmt.Println("Выводим с цветами")
	colorPrint(matrix, resNash, reshPareto)

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
	fmt.Printf("Находим оптимум по Парето\n")
	reshPareto = eqPareto(matrix)
	printCells(matrix, reshPareto)
	cross = crossing(resNash, reshPareto)
	if len(cross) == 0 {
		fmt.Printf("\nПересечений решений нет\n")
	} else {
		fmt.Printf("\nПересечение решений\n")
		printCells(matrix, cross)
	}
	fmt.Println("Выводим с цветами")
	colorPrint(matrix, resNash, reshPareto)

	fmt.Println("\n\nПроверяем функции на игре \"Семейный спор\"")
	matrix, err = getMatrix(disputePath)
	if err != nil {
		panic(err)
	}
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("Находим оптимум по Парето\n")
	reshPareto = eqPareto(matrix)
	printCells(matrix, reshPareto)
	cross = crossing(resNash, reshPareto)
	if len(cross) == 0 {
		fmt.Printf("\nПересечений решений нет\n")
	} else {
		fmt.Printf("\nПересечение решений\n")
		printCells(matrix, cross)
	}
	fmt.Println("Выводим с цветами")
	colorPrint(matrix, resNash, reshPareto)

	fmt.Println("\n\nПроверяем функции на игре \"Дилемма заключённого\"")
	matrix, err = getMatrix(prisonersPath)
	if err != nil {
		panic(err)
	}
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("Находим оптимум по Парето\n")
	reshPareto = eqPareto(matrix)
	printCells(matrix, reshPareto)
	cross = crossing(resNash, reshPareto)
	if len(cross) == 0 {
		fmt.Printf("\nПересечений решений нет\n")
	} else {
		fmt.Printf("\nПересечение решений\n")
		printCells(matrix, cross)
	}
	fmt.Println("Выводим с цветами")
	colorPrint(matrix, resNash, reshPareto)

	fmt.Println("\n\nРешаем игру своего варианта")
	matrix, err = getMatrix(defaultPath)
	if err != nil {
		panic(err)
	}
	printWins(matrix)
	fmt.Printf("Находим равновесие по Нэшу\n")
	resNash = eqNash(matrix)
	printCells(matrix, resNash)
	fmt.Printf("Находим оптимум по Парето\n")
	reshPareto = eqPareto(matrix)
	printCells(matrix, reshPareto)
	cross = crossing(resNash, reshPareto)
	if len(cross) == 0 {
		fmt.Printf("\nПересечений решений нет\n")
	} else {
		fmt.Printf("\nПересечение решений\n")
		printCells(matrix, cross)
	}
	fmt.Println("Выводим с цветами")
	colorPrint(matrix, resNash, reshPareto)
	x, y, v1, v2 := solution(matrix)
	fmt.Printf("Стратегия игрока A ")
	printSlice(x)
	fmt.Printf("Стратегия игрока B ")
	printSlice(y)
	fmt.Printf("Выигрыш первого %.3f\nВыигрыш второго %.3f\n", v1, v2)

}
