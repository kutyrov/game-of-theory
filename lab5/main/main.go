package main

import (
	"fmt"
)

type coalition struct {
	members []int
	value   int
}

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	number := 1
	for i := 1; i <= n; i++ {
		number *= i
	}
	return number
}

func intersection(a, b coalition) []int {
	data := make([]int, 0)
	for _, i := range a.members {
		for _, j := range b.members {
			if i == j {
				data = append(data, i)
			}
		}
	}
	return data
}

func union(a, b coalition) []int {
	i, j := 0, 0
	data := make([]int, 0)
	for i < len(a.members) || j < len(b.members) {
		if a.members[i] < b.members[j] {
			data = append(data, a.members[i])
			i += 1
		} else if a.members[i] > b.members[j] {
			data = append(data, b.members[j])
			j += 1
		} else {
			data = append(data, a.members[i])
			i += 1
			j += 1
		}
	}
	return data
}

func binToNumbers(data []int) []int {
	numbers := make([]int, 0)
	for i := range data {
		if data[i] == 1 {
			numbers = append(numbers, i+1)
		}
	}
	return numbers
}

func binToLen(data []int, l int) []int {
	if l <= len(data) {
		return data
	}
	newData := make([]int, l)
	copy(newData, data)
	return newData
}

func intToBin(number int) []int {
	data := make([]int, 1)
	if number <= 0 {
		return data
	}
	n := 1
	degree := 0
	for n*2 <= number {
		n *= 2
		degree += 1
		data = append(data, 0)
	}
	for number > 0 {
		if n <= number {
			number = number - n
			data[degree] = 1
		}
		n /= 2
		degree -= 1
	}
	return data
}

func main() {
	//d1 := intToBin(5)
	//fmt.Println(binToLen(d1, 4))
	//fmt.Println(binToNumbers(d1))
	//d2 := intToBin(6)
	//fmt.Println(binToLen(d1, 4))
	//fmt.Println(binToNumbers(d2))
	//c1 := coalition{binToNumbers(d1), 5}
	//c2 := coalition{binToNumbers(d2), 6}
	//fmt.Println(intersection(c1,c2))
	//fmt.Println(union(c1,c2))
	fmt.Println(factorial(3))
}
