package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Step struct {
	AChoice        int
	BChoice        int
	AWin           []float64
	BLose          []float64
	EvaluationUp   float64
	EvaluationDown float64
	Epsilon        float64
}

func NewStep() Step {
	s := Step{
		AChoice:        0,
		BChoice:        0,
		AWin:           make([]float64, 0),
		BLose:          make([]float64, 0),
		EvaluationUp:   0.0,
		EvaluationDown: 0.0,
		Epsilon:        0.0,
	}
	return s
}

// при равных выигрышах функция выберет случаную стратегию
func RandChoiceA(data []float64) int {
	max, _ := MaxFromSlice(data)
	indexes := make([]int, 0)
	for index := range data {
		if data[index] == max {
			indexes = append(indexes, index)
		}
	}
	randomIndex := RandomNumber(len(indexes))
	return indexes[randomIndex]
}

// при равных проигрышах функция выберет случайную стратегию
func RandChoiceB(data []float64) int {
	min, _ := MinFromSlice(data)
	indexes := make([]int, 0)
	for index := range data {
		if data[index] == min {
			indexes = append(indexes, index)
		}
	}
	randomIndex := RandomNumber(len(indexes))
	return indexes[randomIndex]
}

// более-менее красивый вывод таблицы шагов
func PrintSteps(steps []Step) {
	fmt.Printf("	A	B	")
	for index := range steps[0].AWin {
		fmt.Printf("x%d		", index+1)
	}
	for index := range steps[0].BLose {
		fmt.Printf("y%d		", index+1)
	}
	fmt.Printf("EvUp	EvDn	Eps\n")
	for index := range steps {
		fmt.Printf("N%d	%d	%d	", index+1, steps[index].AChoice+1, steps[index].BChoice+1)
		for jindex := range steps[index].AWin {
			fmt.Printf("%.2f		", steps[index].AWin[jindex])
		}
		for jindex := range steps[index].BLose {
			fmt.Printf("%.2f		", steps[index].BLose[jindex])
		}
		fmt.Printf("%.2f	%.2f	%.2f\n",
			steps[index].EvaluationUp,
			steps[index].EvaluationDown,
			steps[index].Epsilon,
		)
	}
}

// функия рассичитывает погрешность эпсилон по всей таблице шагов
func CalculateEpsion(steps []Step) float64 {
	maxs := make([]float64, len(steps))
	mins := make([]float64, len(steps))
	for index := range steps {
		maxs[index] = steps[index].EvaluationUp
		mins[index] = steps[index].EvaluationDown
	}
	minmax, _ := MinFromSlice(maxs)
	maxmin, _ := MaxFromSlice(mins)
	return minmax - maxmin
}

// функция находит максимальное число в слайсе
func MaxFromSlice(s []float64) (float64, int) {
	if len(s) == 0 {
		return 0, 0
	}
	max := s[0]
	index := 0
	for i := range s {
		if s[i] > max {
			max = s[i]
			index = i
		}
	}
	return max, index
}

// функция находит минимальное число в слайсе
func MinFromSlice(s []float64) (float64, int) {
	if len(s) == 0 {
		return 0, 0
	}
	min := s[0]
	index := 0
	for i := range s {
		if s[i] < min {
			min = s[i]
			index = i
		}
	}
	return min, index
}

// функция генерирует случайное число в интервале [0,data)
func RandomNumber(data int) int {
	if data <= 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(data)
}

// функция пытается найти седловую точку матрицы и возвращает ошибку если её нет
func SaddlePoint(matrix [][]float64) (int, int, float64, error) {
	maxs := make([]float64, len(matrix[0]), len(matrix[0]))
	mins := make([]float64, len(matrix), len(matrix))

	// найдем минимумы по строкам
	for index := range matrix {
		mins[index] = matrix[index][0]
		for jindex := range matrix[0] {
			if matrix[index][jindex] < mins[index] {
				mins[index] = matrix[index][jindex]
			}
		}
	}

	// найдем максимумы по столбцам
	for jindex := range matrix[0] {
		maxs[jindex] = matrix[0][jindex]
		for index := range matrix {
			if matrix[index][jindex] > maxs[jindex] {
				maxs[jindex] = matrix[index][jindex]
			}
		}
	}

	// найдем minmax и maxmin
	minmax := maxs[0]
	minmaxIndex := 0
	for index := range maxs {
		if maxs[index] < minmax {
			minmax = maxs[index]
			minmaxIndex = index
		}
	}

	maxmin := mins[0]
	maxminIndex := 0
	for index := range mins {
		if mins[index] > maxmin {
			maxmin = mins[index]
			maxminIndex = index
		}
	}

	if minmax == maxmin {
		return maxminIndex + 1, minmaxIndex + 1, minmax, nil
	} else {
		return -1, -1, 0, errors.New("Седловой точки нет")
	}
}

func BrownRobinson(matrix [][]float64) ([]float64, []float64, float64) {
	strategics := len(matrix)
	steps := make([]Step, 0)
	stepNumber := 0

	// Делаем первый шаг алгоритма (заполнение первой строки таблицы)
	steps = append(steps, NewStep())
	steps[0].AChoice = RandomNumber(strategics)
	steps[0].BChoice = RandomNumber(strategics)
	for index := 0; index < strategics; index++ {
		steps[0].AWin = append(steps[0].AWin, matrix[index][steps[0].AChoice])
		steps[0].BLose = append(steps[0].BLose, matrix[steps[0].BChoice][index])

	}
	steps[0].EvaluationUp, _ = MaxFromSlice(steps[0].AWin)
	steps[0].EvaluationDown, _ = MinFromSlice(steps[0].BLose)
	steps[0].Epsilon = CalculateEpsion(steps)

	// остальные строки таблицы заполняем в цикле
	// условие остановки соглано заданию по погрешности
	for steps[stepNumber].Epsilon > 0.1 {
		stepNumber += 1
		steps = append(steps, NewStep())
		steps[stepNumber].AChoice = RandChoiceA(steps[stepNumber-1].AWin)
		steps[stepNumber].BChoice = RandChoiceB(steps[stepNumber-1].BLose)
		for index := 0; index < strategics; index++ {
			steps[stepNumber].AWin = append(steps[stepNumber].AWin,
				steps[stepNumber-1].AWin[index]+matrix[index][steps[stepNumber].BChoice])
			steps[stepNumber].BLose = append(steps[stepNumber].BLose,
				steps[stepNumber-1].BLose[index]+matrix[steps[stepNumber].AChoice][index])
		}
		temp, _ := MaxFromSlice(steps[stepNumber].AWin)
		steps[stepNumber].EvaluationUp = temp / (float64(stepNumber) + 1)
		temp, _ = MinFromSlice(steps[stepNumber].BLose)
		steps[stepNumber].EvaluationDown = temp / (float64(stepNumber) + 1)
		steps[stepNumber].Epsilon = CalculateEpsion(steps)

	}

	// Получаем итоговое решение из таблицы
	value := (steps[stepNumber].EvaluationUp + steps[stepNumber].EvaluationDown) / 2
	aStrategy := make([]float64, strategics)
	bStrategy := make([]float64, strategics)
	for index := range steps {
		aStrategy[steps[index].AChoice] += 1
		bStrategy[steps[index].BChoice] += 1
	}
	for index := range aStrategy {
		aStrategy[index] /= float64(stepNumber + 1)
	}
	for index := range bStrategy {
		bStrategy[index] /= float64(stepNumber + 1)
	}

	// PrintSteps(steps)

	return aStrategy, bStrategy, value
}

func ComputeH(a, b, c, d, e, x, y float64) float64 {
	return a*x*x + b*y*y + c*x*y + d*x + e*y
}

func printMatrix(matrix [][]float64) {
	if len(matrix) == 0 {
		return
	}
	for row := range matrix {
		printSlice(matrix[row])
	}
}

func AnalyticalMethod(data []float64) (float64, float64, float64) {

	// проверяем коэффициенты игры
	if !checkTypeOfGame(data[0], data[1]) {
		panic("Игра не выпукло-вогнутая")
	} else {
		fmt.Printf("Данная игра выпукло-вогнутая\n")
	}

	a, b, c, d, e := data[0], data[1], data[2], data[3], data[4]

	y := (c*d - 2*a*e) / (4*b*a - c*c)

	x := -(c*y + d) / (2 * a)

	h := ComputeH(a, b, c, d, e, x, y)

	return x, y, h
}

func ApproximationOnGrid(data []float64) (float64, float64, float64) {
	// переделать условие остановки
	size := 2
	results := make([][]float64, 0)
	index := 0
	counter := 0
	hFinal := 0.0
	for counter < 10 {

		results = append(results, make([]float64, 3))
		matrix := makeMatrixH(data, size)
		if index < 10 {
			fmt.Printf("шаг %d\n", index+1)
			// выводим матрицу H
			fmt.Println("Полученная матрица H")
			printMatrix(matrix)
			//fmt.Println(matrix)
		}
		x, y, h, error := SaddlePoint(matrix)
		xFloat := float64(x-1) / (float64(size) - 1)
		yFloat := float64(y-1) / (float64(size) - 1)
		y /= (size - 1)
		if error == nil {
			if index < 10 {
				fmt.Println("Есть седловая точка:")
				fmt.Printf("x=%.2f y=%.2f h=%.2f\n\n", xFloat, yFloat, h)
			}
			results[index][0] = xFloat
			results[index][1] = yFloat
			results[index][2] = h

		} else {

			aStrategy, bStrategy, h := BrownRobinson(matrix)
			_, indexA := MaxFromSlice(aStrategy)
			_, indexB := MaxFromSlice(bStrategy)
			results[index][0] = float64(indexA) / (float64(size) - 1)
			results[index][1] = float64(indexB) / (float64(size) - 1)
			results[index][2] = h
			if index < 10 {
				fmt.Println("Седловой точки нет. Решаем методом Брауна-Робинсон")
				fmt.Printf("x=%.2f y=%.2f h=%.2f\n\n",
					results[index][0],
					results[index][1],
					results[index][2],
				)
			}
		}
		if hFinal == 0 {
			hFinal = results[index][2]
		} else if Abs(hFinal-results[index][2]) < 0.01 {
			counter += 1
		} else {
			counter = 0
			hFinal = results[index][2]
		}
		if counter == 10 {
			fmt.Printf("Ответ получен на шаге %d\n", index+1)
			fmt.Printf("x=%.2f y=%.2f h=%.2f\n\n",
				results[index][0],
				results[index][1],
				results[index][2],
			)
		}
		//fmt.Printf("counter = %d 	h = %.2f	hfinal = %.2f", counter, results[index][2], hFinal)
		//fmt.Println()
		index += 1
		size += 1
	}
	steps := len(results) - 1
	return results[steps][0], results[steps][1], results[steps][2]
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func makeMatrixH(data []float64, size int) [][]float64 {
	if size <= 1 {
		return nil
	}

	matrix := make([][]float64, size)
	for index := range matrix {
		matrix[index] = make([]float64, size)
	}

	for index := range matrix {
		for jindex := range matrix[index] {
			matrix[index][jindex] = ComputeH(
				data[0],
				data[1],
				data[2],
				data[3],
				data[4],
				float64(index)/(float64(size)-1),
				float64(jindex)/(float64(size)-1),
			)
		}
	}
	return matrix

}
