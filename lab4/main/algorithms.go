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
	strategies [][]int
	data       [][]int
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

func calculateValue(n *node) {
	//fmt.Printf("Обрабатывается вершина %p\n", n)
	if n == nil {
		return
	}

	if n.children == nil {
		return
	}

	// если не все потомки являются листьями
	for i := range n.children {

		if len(n.children[i].data) == 0 {
			calculateValue(n.children[i])
		}
	}

	// если текущая вершина это дерево, то уже ничего не делаем
	if n.player == -1 {
		return
	}

	// сначала найдем максмумы внутри одного потомка и сравним с максимами других потомков
	maxs := make([]int, len(n.children))
	maxIndexes := make([]int, 0)
	for i := range maxs {
		maxs[i] = minRand
	}

	// нашли максимумы внутри каждого потомка и записали их в массив
	for i := range n.children {

		for _, data := range n.children[i].data {
			if data[n.player] > maxs[i] {
				maxs[i] = data[n.player]
			}
		}
	}

	// определили список потомков, которые надо перенести выше
	max := minRand
	for i, v := range maxs {
		if v > max {
			maxIndexes = make([]int, 0)
			maxIndexes = append(maxIndexes, i)
			max = v
		} else if v == max {
			maxIndexes = append(maxIndexes, i)
		}
	}

	for _, i := range maxIndexes {
		for j := range n.children[i].data {
			n.data = append(n.data, n.children[i].data[j])
			if len(n.children[i].strategies) != 0 {
				n.strategies = append(n.strategies, appendFirst(n.children[i].strategies[j], i))
			} else {
				temp := make([]int, 1)
				temp[0] = i
				n.strategies = append(n.strategies, temp)
			}
		}
	}

	n.children = nil
	t := n
	for t.parent != nil {
		t = t.parent
	}
	printTree(t)
	return
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
	head := node{nil, -1, nil, make([][]int, 0), make([][]int, 0)}
	head.children = make([]*node, len(strategies))

	// на первом уровне должны присутстовать все игроки
	for i := range strategies {
		head.children[i] = &node{&head, i, nil, make([][]int, 0), make([][]int, 0)}
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
				n.children = nil
				n.data = make([][]int, 0)
			}
			for _, n := range thisLevel {
				// fmt.Printf("%p ", n)
				// fmt.Println(n)
				n.children = make([]*node, strategies[n.player])

				if randBool() {
					for index := 0; index < strategies[n.player]; index++ {
						n.children[index] = &node{
							n,
							(n.player + 1) % len(strategies),
							nil,
							make([][]int, 0),
							make([][]int, 0),
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
					n.data = append(n.data, randSlice(len(strategies)))
					//n.data =
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
				leaf.data = append(leaf.data, randSlice(len(strategies)))
				//leaf.data = randSlice(len(strategies))
			}
		}
	}
	return &head
}
