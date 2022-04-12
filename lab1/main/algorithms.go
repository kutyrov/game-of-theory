package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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
	max := MaxFromSlice(data)
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
	min := MinFromSlice(data)
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
	minmax := MinFromSlice(maxs)
	maxmin := MaxFromSlice(mins)
	return minmax - maxmin
}

// функция находит максимальное число в слайсе
func MaxFromSlice(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// функция находит минимальное число в слайсе
func MinFromSlice(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
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

func sliceToPoints(data []float64) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = data[i]
	}
	return pts
}

func drawPlot(steps []Step) error {
	rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = "lab 1 results"
	p.X.Label.Text = "Steps"
	p.Y.Label.Text = "Value"

	data := make([]float64, len(steps))
	for i := range steps {
		data[i] = steps[i].Epsilon
	}

	err := plotutil.AddLinePoints(p,
		"Epsilon", sliceToPoints(data),
	)
	if err != nil {
		return err
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		return err
	}
	return nil
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
	steps[0].EvaluationUp = MaxFromSlice(steps[0].AWin)
	steps[0].EvaluationDown = MinFromSlice(steps[0].BLose)
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
		steps[stepNumber].EvaluationUp =
			MaxFromSlice(steps[stepNumber].AWin) / (float64(stepNumber) + 1)
		steps[stepNumber].EvaluationDown =
			MinFromSlice(steps[stepNumber].BLose) / (float64(stepNumber) + 1)
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

	PrintSteps(steps)

	if err := drawPlot(steps); err != nil {
		panic(err)
	}

	return aStrategy, bStrategy, value
}

func AnalyticalMethod(matrix [][]float64) ([]float64, []float64, float64, error) {
	if len(matrix) != len(matrix[0]) {
		return nil, nil, 0, errors.New("Подана неправильная матрица")
	}

	x := make([]float64, len(matrix), len(matrix))
	y := make([]float64, len(matrix), len(matrix))

	// вычислим одинаковый для x, y и v знаменатель
	denomiator := 0.0
	for index := range matrix {
		for jindex := range matrix {
			denomiator += matrix[index][jindex]
		}
	}

	// вычисляем x* и y*
	for index := range matrix {
		for jindex := range matrix {
			x[index] += matrix[jindex][index]
			y[index] += matrix[index][jindex]
		}
		x[index] /= denomiator
		y[index] /= denomiator
	}

	// вычисляем цену игры v
	var v float64 = 1 / denomiator
	return x, y, v, nil

}
