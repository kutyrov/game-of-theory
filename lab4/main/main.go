package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	minRand     = -5
	maxRand     = 25
	defaultPath = "data/12.txt"
)

func main() {
	height := 0
	players := 0
	strategies := make([]int, 0)
	// считываем данные с файла
	file, err := os.Open(defaultPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	index := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if index == 0 {
			height, _ = strconv.Atoi(numbers[0])
			players, _ = strconv.Atoi(numbers[1])
		} else if index == 1 {
			for i := 0; i < players; i++ {
				number, _ := strconv.Atoi(numbers[i])
				strategies = append(strategies, number)
			}
		} else {
			panic("Неправильный формат файла")
		}
		index += 1
	}
	// fmt.Println(height, players, strategies)
	head := generateTree(height, strategies)
	//fmt.Println(*head)
	printTree(head)
}
