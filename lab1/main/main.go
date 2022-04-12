package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//const defaultPath = "data/example.txt"

const defaultPath = "data/12.txt"

//const defaultPath = "data/temp2.txt"

// удобный вывод слайса
func printSlice(data []float64) {
	if len(data) == 0 {
		return
	}
	fmt.Printf("%.2f", data[0])
	for index := 1; index < len(data); index++ {
		fmt.Printf(" %.2f", data[index])
	}
	fmt.Printf("\n")
}

// функция считывания размеров матрицы из файла
func getSizes(path string) (int, int, error) {

	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows, columns := 0, 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 2 {
			return 0, 0, errors.New("неправильное число аргументов в первой строке")
		}

		rows, err = strconv.Atoi(numbers[0])
		if err != nil {
			return 0, 0, errors.New("первый аргумент это не число")
		}

		columns, err = strconv.Atoi(numbers[1])
		if err != nil {
			return 0, 0, errors.New("второй аргумент это не число")
		}

		if rows <= 0 || columns <= 0 {
			return 0, 0, errors.New("размер матрицы не может быть отрицательным")
		}
		break
	}

	return rows, columns, nil

}

// функция считывания матрицы из файла
// возможно стоит сделать композицию с getSizes
func getMatrix(path string, rows, columns int) ([][]float64, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	index := -1
	matrix := make([][]float64, rows, rows)
	for jindex := range matrix {
		matrix[jindex] = make([]float64, columns, columns)
	}

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		index += 1

		if index == 0 {
			continue
		}

		if index > rows {
			return nil, errors.New("Слишком много строк")
		}
		if len(numbers) != columns {
			return nil, errors.New("Неправильное число столбцов в матрице")
		}
		for jindex, number := range numbers {
			temp, err := strconv.Atoi(number)
			matrix[index-1][jindex] = float64(temp)
			if err != nil {
				return nil, err
			}
		}
	}
	if index < rows {
		return nil, errors.New("неправильное число строк")
	}
	return matrix, nil
}

func main() {

	// считываем данные
	path := ""
	fmt.Println("Введите имя файла или нажмите enter")
	fmt.Scanf("%s\n", &path)
	if path == "" {
		path = defaultPath
	}

	rows, columns, error := getSizes(path)
	if error != nil {
		panic(error)
	}

	matrix, error := getMatrix(path, rows, columns)
	if error != nil {
		panic(error)
	}

	fmt.Printf("Размеры матрицы %d на %d\n", rows, columns)
	fmt.Println(rows, columns)
	fmt.Println("\nСама матрица:")
	for index := range matrix {
		for jindex := range matrix[index] {
			fmt.Printf("%d ", int(matrix[index][jindex]))
		}
		fmt.Println()
	}

	// проверим наличие седловой точки
	aDominant, bDominant, value, err := SaddlePoint(matrix)
	if err == nil {
		fmt.Printf("Седловая точка есть. Игрок А должен использовать стратегию %d,"+
			"игрок B должен использовать стратегию %d, цена игры равна %.2f\n",
			aDominant,
			bDominant,
			value,
		)
	} else {
		fmt.Println("Седловой точки нет")
	}

	// Аналитичесий метод
	fmt.Println("\nРешаем аналитическим методом")
	C, err := InverseMatrix(matrix)
	if err != nil {
		panic(err)
	}

	x, y, v, _ := AnalyticalMethod(C)

	fmt.Printf("Стратегия игрока А: ")
	printSlice(x)
	fmt.Printf("Стратегия игрорка B: ")
	printSlice(y)
	fmt.Printf("Цена игры равна %.2f\n", v)

	// Алгоритм Брауна-Робинсон
	fmt.Println("\nРешаем методом Брауна-Робинсон")
	aStrategy, bStrategy, value := BrownRobinson(matrix)
	fmt.Printf("Стратегия игрока А: ")
	printSlice(aStrategy)
	fmt.Printf("Стратегия игрорка B: ")
	printSlice(bStrategy)
	fmt.Printf("Цена игры равна %.2f\n", value)

}
