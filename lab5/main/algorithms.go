package main

import (
	"fmt"
	"math"
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

// возвращает пересечение членов коалиции
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

// возвращает множество a\b
func difference(a, b []int) []int {
	//fmt.Println("получили", a, b)
	i, j := 0, 0
	data := make([]int, 0)
	for i < len(a) && j < len(b) {
		//fmt.Println(a.members[i], b.members[j], data)
		if a[i] < b[j] {
			data = append(data, a[i])
			i += 1
		} else if a[i] > b[j] {
			j += 1
		} else {
			i += 1
			j += 1
		}
	}
	if i < len(a) {
		for ; i < len(a); i++ {
			data = append(data, a[i])
		}
	}
	//fmt.Println("вернули", data)
	return data
}

// возвращает объединение членов коалиции
func union(a, b coalition) []int {
	i, j := 0, 0
	data := make([]int, 0)
	// fmt.Println()
	// fmt.Println("hello")
	// fmt.Println(a, b)
	for i < len(a.members) && j < len(b.members) {
		//fmt.Println(a.members[i], b.members[j], data)
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
	if i < len(a.members) {
		for ; i < len(a.members); i++ {
			data = append(data, a.members[i])
		}
	}
	if j < len(b.members) {
		for ; j < len(b.members); j++ {
			data = append(data, b.members[j])
		}
	}
	// fmt.Println(data)
	// fmt.Println("hello")
	// fmt.Println()
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

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// возвращает коалицию по списку входящих в неё членов
func findCoalition(data []coalition, members []int) coalition {
	for _, c := range data {
		if isEqual(c.members, members) {
			return c
		}
	}
	return coalition{}
}

func intToNumbers(data int) []int {
	return binToNumbers(intToBin(data))
}

func isSuperaddity(data []coalition) bool {
	flag := true
	fmt.Println("Формат вывода S + T >=< S u T")
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if len(intersection(data[i], data[j])) == 0 {
				u := findCoalition(data, union(data[i], data[j]))
				if u.value < (data[i].value + data[j].value) {
					flag = false
					fmt.Printf("v%d + v%d > v%d\n",
						data[i].members,
						data[j].members,
						u.members,
					)
					fmt.Printf("%d + %d > %d\n\n",
						data[i].value,
						data[j].value,
						u.value,
					)
				} else {
					fmt.Printf("v%d + v%d <= v%d\n",
						data[i].members,
						data[j].members,
						u.members,
					)
					fmt.Printf("%d + %d <= %d\n\n",
						data[i].value,
						data[j].value,
						u.value,
					)

				}
				//fmt.Println(data[i], data[j], u)
			}

		}
	}
	return flag
}

func isConvex(data []coalition) bool {
	fmt.Println("Формат вывода (S u T) + (S n T) >=< S + T")
	flag := true
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			inter := findCoalition(data, intersection(data[i], data[j]))
			uni := findCoalition(data, union(data[i], data[j]))
			if (inter.value + uni.value) < (data[i].value + data[j].value) {
				flag = false
				fmt.Printf("v%d + v%d < v%d + v%d\n",
					inter.members,
					uni.members,
					data[i].members,
					data[j].members,
				)
				fmt.Printf("%d + %d < %d + v%d\n\n",
					inter.value,
					uni.value,
					data[i].value,
					data[j].value,
				)
			} else {
				fmt.Printf("v%d + v%d >= v%d + v%d\n",
					inter.members,
					uni.members,
					data[i].members,
					data[j].members,
				)
				fmt.Printf("%d + %d >= %d + %d\n\n",
					inter.value,
					uni.value,
					data[i].value,
					data[j].value,
				)

			}
		}
	}
	return flag
}

func printSlice(data []float64) {
	if len(data) == 0 {
		return
	}
	fmt.Printf("%.2f", data[0])
	for index := 1; index < len(data); index++ {
		fmt.Printf("	%.2f", data[index])
	}
	fmt.Printf("\n")
}

func isContain(data []int, v int) bool {
	for _, i := range data {
		if i == v {
			return true
		}
	}
	return false
}

func calculateCoefficients(players int, data []coalition) []float64 {
	coefs := make([]float64, players)
	for i := 0; i < players; i++ {
		n := 0
		for j := range data {
			S := data[j]
			if isContain(S.members, i+1) {
				t := make([]int, 1)
				t[0] = i + 1
				T := findCoalition(data, difference(S.members, t))
				//fmt.Println(S, t, T)
				n += factorial(len(S.members)-1) *
					factorial(players-len(S.members)) *
					(S.value - T.value)
			}
		}
		//fmt.Println("")
		//n += factorial(len(S.members) - 1) * data.value
		coefs[i] = float64(n) / float64(factorial(players))
	}

	// for i := range coefs {
	// 	coefs[i] /=
	// }
	return coefs
}

func firstAksiom(coefs []float64, data []coalition, players int) {
	fmt.Println("\nГрупповая рационализация:")
	sum := 0.0
	for i := range coefs {
		sum += coefs[i]
	}
	lastIndex := int(math.Pow(2, float64(players)) - 2)
	fmt.Printf("Cумма коэффициентов равна %.2f,а v(I) = %d\n", sum, data[lastIndex].value)
}

func secondAksiom(coefs []float64, data []coalition, players int) {
	fmt.Println("\nИндивидуальная рационализация:")
	for i := 0; i < players; i++ {
		temp := make([]int, 1)
		temp[0] = i + 1
		c := findCoalition(data, temp)
		fmt.Printf("x%d = %.2f >= v({%d}) = %d\n", i+1, coefs[i], i+1, c.value)
	}
}
