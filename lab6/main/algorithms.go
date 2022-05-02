package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

//генерирует случайное целое в интервале [min,max]
func randInt(min, max int) int {
	if max < min {
		return 0
	}
	return rand.Intn(max-min+1) + min
}

// возвращает слайс длины n с числами
func randomSlice(n, maxRand int) []float64 {
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = float64(rand.Intn(maxRand))
	}
	return data
}

func printSlice(data []float64) {
	if len(data) == 0 {
		return
	}
	fmt.Printf("%.3f", data[0])
	for index := 1; index < len(data); index++ {
		fmt.Printf("\t%.3f", data[index])
	}
	fmt.Printf("\n")
}

func printIntSlice(data []int) {
	if len(data) == 0 {
		return
	}
	fmt.Printf("%d", data[0])
	for index := 1; index < len(data); index++ {
		fmt.Printf("\t%d", data[index])
	}
	fmt.Printf("\n")
}

func normSlice(data []float64) []float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}

	nData := make([]float64, len(data))
	for i := range data {
		nData[i] = data[i] / float64(sum)
	}
	return nData
}

func arithmeticMean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	mean := 0.0
	for _, v := range data {
		mean += v
	}
	mean /= float64(len(data))
	return mean
}

func abs(v float64) float64 {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}

func epsilon(data []float64) float64 {
	mean := arithmeticMean(data)
	eps := 0.0
	for _, v := range data {
		eps += abs(mean - v)
	}
	eps /= float64(len(data))
	return eps
}

// возвращает число шагов до достижения eps и итоговый вектор
func stepFirst(matrix [][]float64, x []float64) (int, []float64) {
	steps := 0
	for epsilon(x) > maxEpsilon {
		steps++
		temp := transposeMatrix(sliceToMatrix(x))
		x = matrixToSlice(transposeMatrix(multMatrix(matrix, temp)))

		//printSlice(x)
	}

	return steps, x
}

// возвращает число шагов и итоговый вектор мнений агентов
func stepSecond(matrix [][]float64, players int) (int, float64) {
	steps := 0
	if players < 2 {
		panic("Мало игроков")
	}
	// сформируем вектор с учётов того, что играет players игроков
	x := make([]float64, len(matrix))
	firstWin := float64(randInt(0, firstPlayerWin))
	secondWin := float64(randInt(secondPlayerWin, 0))
	agents := make([]int, len(matrix))

	for i := range x {
		agent := rand.Intn(players + 1)
		switch agent {
		// 0 - никому не принадлежащий агент
		case 0:
			x[i] = float64(randInt(secondPlayerWin, firstPlayerWin))
		// 1 - агент первого игрока
		case 1:
			x[i] = firstWin
		// 2 - агент второго игрока
		case 2:
			x[i] = secondWin
		}
		agents[i] = agent
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprint(w, "Вектор мнений и распределение агентов:\n")
	for _, v := range x {
		fmt.Fprintf(w, "\t%.3f", v)
	}
	fmt.Fprint(w, "\t\n")
	for _, v := range agents {
		fmt.Fprintf(w, "\t%d", v)
	}
	fmt.Fprint(w, "\t\n")

	for epsilon(x) > maxEpsilon {
		steps++
		temp := transposeMatrix(sliceToMatrix(x))
		x = matrixToSlice(transposeMatrix(multMatrix(matrix, temp)))
	}

	fmt.Fprint(w, "\nИтоговый вектор мнений\n")
	for _, v := range x {
		fmt.Fprintf(w, "\t%.3f", v)
	}
	fmt.Fprint(w, "\t\n")
	w.Flush()

	if x[0] > 0 {
		fmt.Println("Победил первый игрок")
	} else {
		fmt.Println("Победил второй игрок")
	}
	return steps, x[0]
}
