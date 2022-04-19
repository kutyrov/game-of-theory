package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	parent     *node
	player     int
	children   []*node
	strategies []int
	data       []int
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

func appendFirst(data []int, value int) []int {
	temp := make([]int, 0)
	temp = append(temp, value)
	for _, i := range data {
		temp = append(temp, i)
	}
	return temp
}

func calculateValue(n *node) []int {
	//fmt.Printf("Обрабатывается вершина %p\n", n)
	if n == nil {
		return nil
	}

	// если не все потомки являются листьями
	for i := range n.children {
		if n.children[i] == nil {
			continue
		}
		if n.children[i].children != nil {
			//fmt.Printf("уход по стеку от вершины %p\n", n.children[i])
			calculateValue(n.children[i])
		}
	}

	// если все потомки являются листьями
	if n.player == -1 {
		return nil
	}
	max := minRand - 1
	maxIndex := -1
	for i := range n.children {
		if n.children[i] != nil {
			if n.children[i].data[n.player] > max {
				max = n.children[i].data[n.player]
				maxIndex = i
			}
		}
	}
	if maxIndex != -1 {
		n.data = n.children[maxIndex].data
		n.strategies = appendFirst(n.children[maxIndex].strategies, maxIndex)
	}
	n.children = nil
	t := n
	for t.parent != nil {
		t = t.parent
	}
	printTree(t)
	return n.data
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
	// fmt.Printf("%p ", n)
	//fmt.Print(*n)
	//fmt.Println(n.player, n.data)
	fmt.Printf("(%d,%d,%d)\n", n.player, n.data, n.strategies)
	// fmt.Print(n.data)
	// fmt.Printf(")\n")
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
	fmt.Println("Выводим дерево")
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
	head := node{nil, -1, nil, make([]int, 0), nil}
	head.children = make([]*node, len(strategies))

	// на первом уровне должны присутстовать все игроки
	for i := range strategies {
		head.children[i] = &node{&head, i, nil, make([]int, 0), nil}
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
							make([]int, 0),
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
	return &head
}
