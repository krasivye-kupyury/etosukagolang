package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, maxsvz, progres, kolsvz, energi, k, ii int
	var all, s, char string
	var t = false
	energi = 0
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
				all += strconv.Itoa(svz[row][col])
			}
			if col == 1 {
				fmt.Println("введите его знакомого")
				fmt.Scan(&svz[row][col])
				all += strconv.Itoa(svz[row][col])
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
		for m := 1; m <= maxsvz; m++ {
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
			char = string([]rune(s)[ii])
			k, _ = strconv.Atoi(char)
		}
		break
	}
	fmt.Println(s)
	//fmt.Println(all)

}
