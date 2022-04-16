package main

import (
	"fmt"
	"os"
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
	// for index := range matrix {
	// 	if matrix[index][j].AWin >= cell.AWin &&
	// 		matrix[index][j].BWin >= cell.BWin &&
	// 		index != i {
	// 		flag = false
	// 		break
	// 	}
	// 	if matrix[i][index].AWin >= cell.AWin &&
	// 		matrix[i][index].BWin >= cell.BWin &&
	// 		index != j {
	// 		flag = false
	// 		break
	// 	}
	// }
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
	// for row := range matrix {
	// 	fmt.Printf("%d,%d", matrix[row][0].AWin, matrix[row][0].BWin)
	// 	for col := 1; col < len(matrix[row]); col++ {
	// 		fmt.Printf("\t%d,%d", matrix[row][col].AWin, matrix[row][col].BWin)
	// 	}
	// 	fmt.Printf("\n")
	// }

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
