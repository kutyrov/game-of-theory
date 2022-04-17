package main

import (
	"math/rand"
	"time"
)

type node struct {
	parent   int
	children []*node
	data     []int
}

//генерирует случайное целое в интервале [min,max]
func randInt(min, max int) int {
	if max < min {
		return 0
	}
	return rand.Intn(max-min+1) + min
}

func randSlice(n int) []int {
	if n < 0 {
		return nil
	}
	data := make([]int, n)
	for i := range data {
		data[i] = randInt(minRand, maxRand)
	}
	return data
}

func generateTree(height, players int, strategies []int) node {
	rand.Seed(time.Now().UnixNano())
	return node{}
}
