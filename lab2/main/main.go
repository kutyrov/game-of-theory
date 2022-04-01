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

// получаем коэффициенты из файла
func getInput(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]float64, 5, 5)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 5 {
			return nil, errors.New("ожидалось 5 аргументов на входе")
		}

		for index := range numbers {
			data[index], _ = strconv.ParseFloat(numbers[index], 64)
		}

	}
	return data, nil
}

func checkTypeOfGame(a, b float64) bool {
	if 2*a < 0 && 2*b > 0 {
		return true
	}
	return false
}

func main() {

	// считываем данные
	path := ""
	fmt.Println("Введите имя файла или нажмите enter")
	fmt.Scanf("%s\n", &path)
	if path == "" {
		path = defaultPath
	}

	data, err := getInput(path)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Считали из файла следующие коэффициенты:\n")
	printSlice(data)
	fmt.Println()

	fmt.Println("Решаем аналитическим методом:")
	x, y, h := AnalyticalMethod(data)
	fmt.Printf("x=%.2f y=%.2f h=%.2f\n", x, y, h)

	fmt.Printf("\nРешаем методом аппроксимации функции выигрышей на сетке:\n")
	ApproximationOnGrid(data)

}
