package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type pair struct {
	first  int
	second int
}

// функция находит максимальное число в слайсе
func MaxFromSlice(s []int) int {
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

func pointInSlice(data []pair, value pair) bool {
	for _, d := range data {
		if d == value {
			return true
		}
	}
	return false
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}

func chooseColor(f1, f2 bool) string {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"
	colorWhite := "\033[37m"
	if f1 && f2 {
		return colorRed
	} else if f1 {
		return colorGreen
	} else if f2 {
		return colorBlue
	} else {
		return colorWhite
	}
}

func colorPrint(matrix [][]Win, reshNash []pair, reshPareto []pair) {

	w := tabwriter.NewWriter(os.Stdout, 0, 25, 2, ' ', tabwriter.AlignRight)

	for i := range matrix {
		for j := range matrix[i] {
			temp := pair{i, j}
			flagNash := pointInSlice(reshNash, temp)
			flagPareto := pointInSlice(reshPareto, temp)
			out := chooseColor(flagNash, flagPareto) +
				"(" + FloatToString(matrix[i][j].AWin) + "," +
				FloatToString(matrix[i][j].BWin) + ")"
			fmt.Fprint(w, "\t", out)
		}
		fmt.Fprint(w, "\t\n")
	}
	w.Flush()
	fmt.Fprint(w, chooseColor(false, false))
	w.Flush()
}

func eqNash(matrix [][]Win) []pair {
	if len(matrix) != len(matrix[0]) {
		return nil
	}
	rowsMax := make([]int, len(matrix))
	colsMax := make([]int, len(matrix))

	for index := range matrix {
		rowMax := matrix[index][0].BWin
		rowMaxIndex := 0
		colMax := matrix[0][index].AWin
		colMaxIndex := 0
		for jindex := range matrix {
			if matrix[index][jindex].BWin > rowMax {
				rowMax = matrix[index][jindex].BWin
				rowMaxIndex = jindex
			}
			if matrix[jindex][index].AWin > colMax {
				colMax = matrix[jindex][index].AWin
				colMaxIndex = jindex
			}
		}
		rowsMax[index] = rowMaxIndex
		colsMax[index] = colMaxIndex
	}

	strategies := make([]pair, 0)
	for index := range rowsMax {
		if colsMax[rowsMax[index]] == index {
			strategies = append(strategies, pair{index, rowsMax[index]})
		}
	}

	return strategies
}

func checkPoint(matrix [][]Win, i, j int) bool {
	cell := matrix[i][j]
	flag := true
	for index, rows := range matrix {
		for jindex, data := range rows {
			if data.AWin >= cell.AWin &&
				data.BWin >= cell.BWin &&
				(i != index || jindex != j) {
				flag = false
			}
		}
	}
	return flag
}

func eqPareto(matrix [][]Win) []pair {
	strategies := make([]pair, 0)
	for index := range matrix {
		for jindex := range matrix {
			if checkPoint(matrix, index, jindex) {
				strategies = append(strategies, pair{index, jindex})
			}
		}
	}

	return strategies
}

func printSlice(data []float64) {
	if len(data) == 0 {
		return
	}
	fmt.Printf("%.2f", data[0])
	for index := 1; index < len(data); index++ {
		fmt.Printf("\t%.2f", data[index])
	}
	fmt.Printf("\n")
}

func printMatrix(matrix [][]float64) {
	if len(matrix) == 0 {
		return
	}
	for row := range matrix {
		printSlice(matrix[row])
	}
}

func printWins(matrix [][]Win) {
	if len(matrix) == 0 {
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	for row := range matrix {
		//fmt.Fprintf(w, "(%.2f,%.2f)", matrix[row][0].AWin, matrix[row][0].BWin)
		fmt.Fprint(w, "(", matrix[row][0].AWin, ",", matrix[row][0].BWin, ")")
		for col := 1; col < len(matrix[row]); col++ {
			//fmt.Fprintf(w, "\t(%.2f,%.2f)", matrix[row][col].AWin, matrix[row][col].BWin)
			fmt.Fprint(w, "\t(", matrix[row][col].AWin, ",", matrix[row][col].BWin, ")")
		}
		fmt.Fprint(w, "\t\n")
	}
	w.Flush()
}

func solution(matrix [][]Win) ([]float64, []float64, float64, float64) {
	if len(eqNash(matrix)) == 2 {
		A, B := winToMatrix(matrix)
		B, _ = InverseMatrix(B)
		A, _ = InverseMatrix(A)
		v1 := 1 / sumMatrix(A)
		v2 := 1 / sumMatrix(B)

		temp := make([][]float64, 1)
		temp[0] = generateU(len(B))
		x := matrixOnFloat(multMatrix(temp, B), v2)
		//fmt.Println(multMatrix(temp, B), x)

		temp = TransposeMatrix(temp)
		y := TransposeMatrix(matrixOnFloat(multMatrix(A, temp), v1))
		return matrixToSlice(x), matrixToSlice(y), v1, v2
	} else {
		return nil, nil, 0, 0
	}
}
