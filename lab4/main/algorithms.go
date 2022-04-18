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
	//fmt.Println(data)
	return data
}

func printNode(n *node, level int, sep map[int]bool) {
	if n.parent == nil {
		// fmt.Print("корень ", n)
		// fmt.Printf("%p\n", n)
		return
	}
	//fmt.Println(sep)
	for i := 0; i < level-1; i++ {
		v, ok := sep[i]
		if ok {
			if v {
				fmt.Print("|   ")
			} else {
				fmt.Print("    ")
			}
		} else {
			fmt.Print("    ")
		}

	}
	fmt.Print("|---")
	//fmt.Printf("%p ", n)
	//fmt.Println(*n)
	fmt.Println(n.player, n.data)
}

func bypath(n *node, level int, sep map[int]bool) {
	if n == nil {
		// for i := 0; i < level-1; i++ {
		// 	v, ok := sep[i]
		// 	if ok {
		// 		if v {
		// 			fmt.Print("|   ")
		// 		} else {
		// 			fmt.Print("    ")
		// 		}
		// 	} else {
		// 		fmt.Print("    ")
		// 	}
		// }
		// fmt.Println("|---", nil)
		return
	}
	printNode(n, level, sep)
	//fmt.Println("обработка", n)
	if n.children != nil {
		level += 1
		for index, child := range n.children {
			flag := false
			for i := index + 1; i < len(n.children); i++ {
				if n.children[i] != nil {
					flag = true
				}
			}
			if flag {
				sep[level-1] = true
			} else {
				sep[level-1] = false
			}
			bypath(child, level, sep)
		}
		for i := 0; i < level-1; i++ {
			v, ok := sep[i]
			if ok {
				if v {
					fmt.Print("|   ")
				} else {
					fmt.Print("    ")
				}
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
	}
}

func printTree(head *node) {
	level := 0
	bypath(head, level, make(map[int]bool))
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

	//осталось сгенерировать height-1 уровень
	thisLevel := head.children
	//fmt.Println(thisLevel)

	for i := 0; i < height-1; i++ {
		nextLevel := make([]*node, 0)
		// fmt.Println("level ", i+1)
		// убедились, что на следующем уровне будут вершины
		for len(nextLevel) == 0 {
			for _, n := range thisLevel {
				// fmt.Printf("%p ", n)
				// fmt.Println(n)
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
				// fmt.Println("нач")
				// fmt.Println(n)
				if isLeaf {
					n.children = nil
					n.data = randSlice(len(strategies))
				}
				// fmt.Println(n)
				// fmt.Println("кон")
			}
		}
		thisLevel = make([]*node, len(nextLevel))
		copy(thisLevel, nextLevel)
		if i == height-2 {
			for _, leaf := range thisLevel {
				//fmt.Println(leaf)
				leaf.data = randSlice(len(strategies))
			}
		}
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
