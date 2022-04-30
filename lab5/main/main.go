package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//const defaultPath = "data/example.txt"

const defaultPath = "data/12.txt"

// возвращает кол-во игроков и выигрыши коалиций
func getInput(path string) (int, []coalition) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNumber := 0
	players := 0
	data := make([]coalition, 0)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if rowNumber == 0 {
			players, err = strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
		} else {
			if len(numbers) != int((math.Pow(2, float64(players)))-1) {
				panic("Неправильное число параметров в строке выигрышей коалиций")
			}
			for i, sValue := range numbers {
				value, err := strconv.Atoi(sValue)
				if err != nil {
					panic(err)
				}
				data = append(data, coalition{intToNumbers(i + 1), value})
			}
		}

		rowNumber += 1

	}
	return players, data
}

func main() {
	players, input := getInput(defaultPath)
	if isSuperaddity(input) {
		fmt.Println("Игра суперадддитивная")
	} else {
		fmt.Println("Игра не суперадддитивная")
	}
	fmt.Println()

	if isConvex(input) {
		fmt.Println("Игра выпуклая")
	} else {
		fmt.Println("Игра не выпуклая")
	}
	fmt.Println("\nВыводим коэффициенты игроков")
	coefs := calculateCoefficients(players, input)
	printSlice(coefs)

	fmt.Println("\nПроверим аксиомы рационализации")
	firstAksiom(coefs, input, players)
	secondAksiom(coefs, input, players)

}
