package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, maxsvz, progres, kolsvz, energi, k, ii, v, min int
	var all, s, char, ss string
	var t = false
	energi = 0
	min = 1000
	fmt.Println("введите количество жителей")
	fmt.Scan(&n)
	e := make([]int, n)
	fmt.Println("введите энергию для каждого жителя")
	for i := range e {
		fmt.Scan(&e[i])
	}
	for i := range e {
		fmt.Print(" ", i, " ")
	}
	fmt.Println()
	for i := range e {
		fmt.Print(" ", e[i], " ")
	}
	fmt.Println()
	fmt.Println("введите количество связей")
	for i := 1; i <= n-1; i++ {
		progres = progres + i
	}
	maxsvz = (n * n) - n - progres
	fmt.Println("максимум связей может быть ", maxsvz)
	fmt.Scan(&kolsvz)
	fmt.Println(kolsvz)
	fmt.Println("введите связи")
	svz := make([][]int, kolsvz)
	for i := 0; i < kolsvz; i++ {
		svz[i] = make([]int, 2)
	}
	for row := 0; row < kolsvz; row++ {
		for col := 0; col < 2; col++ {
			if col == 0 {
				fmt.Println("введите первого человека ")
				fmt.Scan(&svz[row][col])
			}
			if col == 1 {
				fmt.Println("введите его знакомого")
				fmt.Scan(&svz[row][col])
			}
		}
	}
	fmt.Println("человек  :  пара")
	for row := 0; row < kolsvz; row++ {
		for col := 0; col < 2; col++ {
			fmt.Print("    ", svz[row][col], "    ")
		}
		fmt.Println()
	}
	for m := 1; m <= n; m++ {
		for row := 0; row < kolsvz; row++ {
			for col := 0; col < 2; col++ {
				if m == svz[row][col] {
					t = true
				}
			}

		}
		if t == false {
			energi = e[m-1] + energi
		}
		t = false

	}
	fmt.Println("энергия, чтобы позвать людей, которых нет в связях =", energi)
	k = svz[0][0]
	s = strconv.Itoa(k)
	ii = 0
	for {
		for m := 1; m <= n*n; m++ {
			for row := 0; row < kolsvz; row++ {
				if k == svz[row][0] {
					s += strconv.Itoa(svz[row][1])
				}
			}
			for row := 0; row < kolsvz; row++ {
				if k == svz[row][1] {
					s += strconv.Itoa(svz[row][0])
				}
			}
			ii += 1
			//fmt.Println("do char")
			char = string([]rune(s)[ii])
			//fmt.Println("posle char")
			k, _ = strconv.Atoi(char)
		}
		ss = s
		runes := []rune(ss)
		seen := make(map[rune]bool)
		result := make([]rune, 0)
		for _, r := range runes {
			if !seen[r] {
				seen[r] = true
				result = append(result, r)
			}
		}
		s = string(result)
		all += s
		fmt.Println(s)
		for i := 0; i < len(s); i++ {
			char = string([]rune(s)[i])
			k, _ = strconv.Atoi(char)
			k = k - 1
			if min > e[k] {
				min = e[k]

			}
		}
		fmt.Println("min e ", min)
		energi = energi + min
		for row := 0; row < kolsvz; row++ {
			for col := 0; col < 2; col++ {
				v = svz[row][col]
				char = strconv.Itoa(v)
				t = strings.Contains(all, char)
				if t == false {
					k = v
					fmt.Println("eto k ", k)
					break
				}
				if t == false {
					break
				}
			}
			if t == false {
				break
			}
		}
		ii = 0
		fmt.Println("eto t ", t)
		s = strconv.Itoa(k)
		min = 1000
		if t == true {
			break
		}
	}
	fmt.Println("eto energi ", energi)

}
