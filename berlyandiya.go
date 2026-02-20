package main

import (
	"fmt"
)

func main() {
	var n, maxsvz, progres, kolsvz, energi, k, ii, v, min, c int
	var char int
	var t = false
	energi = 0
	c = 0
	min = 1000
	fmt.Println("введите количество жителей")
	fmt.Scan(&n)
	e := make([]int, n)
	a := make([]int, n*n)
	b := make([]int, n*n)
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
	a[0] = k
	ii = 0
	for {
		for m := 1; m < n*n; m++ {
			for row := 0; row < kolsvz; row++ {
				if k == svz[row][0] {
					a[m] = svz[row][1]
					m = m + 1
					if m == n*n {
						break
					}
				}
			}
			for row := 0; row < kolsvz; row++ {
				if k == svz[row][1] {
					a[m] = svz[row][0]
					m = m + 1
					if m == n*n {
						break
					}
				}
			}
			m = m - 1
			ii += 1
			char = a[ii]
			k = char
		}
		seen := make(map[int]bool)
		result := make([]int, 0)
		for _, num := range a {
			seen[num] = true
		}
		for num := range seen {
			result = append(result, num)
		}
		b = append(b, result...)
		for i := 0; i < len(result); i++ {
			char = result[i]
			k = char
			k = k - 1
			if min > e[k] {
				min = e[k]
			}
		}
		fmt.Println("минимум энергии для группы", min)
		energi = energi + min
		for row := 0; row < kolsvz; row++ {
			for col := 0; col < 2; col++ {
				v = svz[row][col]
				char = v
				for iiii := 0; iiii < len(b); iiii++ {
					if char != b[iiii] {
						c = c + 1
					}
					if char == b[iiii] {
						c = 0
						break
					}
				}
				if c == len(b) {
					k = v
					break
				}
			}
			if c == len(b) {
				break
			}
		}
		ii = 0
		for mm := range a {
			a[mm] = 0
		}
		a[0] = k
		min = 1000
		if c == 0 {
			break
		}
	}
	fmt.Println("минимум энергии чтобы позвать всех ", energi)
}
