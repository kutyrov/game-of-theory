package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	parent   *node
	player   int
	children []*node
	value    int
	data     []int
}

//генерирует случайное целое в интервале [min,max]
func randInt(min, max int) int {
	if max < min {
		return 0
	}
	return rand.Intn(max-min+1) + min
}

func randBool() bool {
	number := rand.Intn(2)
	if number == 0 {
		return false
	} else {
		return true
	}
}

//генерирует случайный слайс длины n
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

func printNode(n *node, level int) {
	if n.parent == nil {
		return
	}
	for i := 0; i < level-1; i++ {
		fmt.Print("    ")
	}
	fmt.Print("|---")
	fmt.Println(n.player, n.data)
}

func bypath(n *node, level int) {
	if n == nil {
		return
	}
	printNode(n, level)
	//fmt.Println("обработка", n)
	if n.children != nil {
		level += 1
		for _, child := range n.children {
			bypath(child, level)
		}
	} else {
		printNode(n, level)
	}
}

func printTree(head *node) {
	level := 0
	bypath(head, level)
}

func generateTree(height int, strategies []int) *node {
	// проверяем входные параметры
	if height <= 0 {
		return nil
	}
	for _, value := range strategies {
		if value <= 0 {
			return nil
		}
	}
	rand.Seed(time.Now().UnixNano())
	// обозначаем корень дерева
	head := node{nil, -1, nil, 0, nil}
	head.children = make([]*node, len(strategies))

	// на первом уровне должны присутстовать все игроки
	for i := range strategies {
		head.children[i] = &node{&head, i, nil, 0, nil}
	}
	//fmt.Println(head.children)
	//осталось сгенерировать height-1 уровень
	thisLevel := head.children
	//fmt.Println(thisLevel)

	for i := 0; i < height-1; i++ {
		nextLevel := make([]*node, 0)
		// убедились, что на следующем уровне будут вершины
		for len(nextLevel) == 0 && i != height-2 {
			for _, n := range thisLevel {
				n.children = make([]*node, strategies[n.player])
				for index := 0; index < strategies[n.player]; index++ {
					if randBool() {
						n.children[index] = &node{
							n,
							(n.player + 1) % len(strategies),
							nil,
							0,
							nil,
						}
						nextLevel = append(nextLevel, n.children[index])
					}
				}
				isLeaf := true
				for _, child := range n.children {
					if child != nil {
						isLeaf = false
					}
				}
				if isLeaf {
					n.children = nil
					n.data = randSlice(len(strategies))
				}
			}
		}
		thisLevel = make([]*node, len(nextLevel))
		copy(thisLevel, nextLevel)
	}
	// в коненые вершины записываем случайные выигрыши
	// thisLevel = head.children
	// for i := 0; i < height-1; i++ {
	// 	nextLevel := make([]*node, 0)
	// 	for _, n := range thisLevel {
	// 		if n.children == nil {
	// 			n.data = randSlice(len(strategies))
	// 		} else {
	// 			nextLevel = append(nextLevel, n)
	// 		}
	// 	}
	// 	thisLevel = make([]*node, len(nextLevel))
	// 	copy(thisLevel, nextLevel)
	// }
	return &head
}
