package main

import (
	"fmt"
)

func main() {
	var n, maxsvz, progres, kolsvz, energi, k, ii, v, min, c, aa, bb, kk int
	var char int
	var t = false
	energi = 0
	kk = 0
	c = 0
	min = 1000
	for {
		fmt.Println("введите количество жителей")
		fmt.Scan(&n)
		n, t = proverkaa(n, t)
		if t == true {
			break
		}
	}
	e := make([]int, n)
	a := make([]int, n*n)
	b := make([]int, n*n)
	t = false
	for {
		fmt.Println("введите энергию для каждого жителя")
		for i := 0; i < len(e); i++ {
			fmt.Scan(&e[i])
			e[i], t = proverkaa(e[i], t)
			if t == false {
				if i == 0 {
					i = i - 1
				}
				if i > 0 {
					i = i - 1
				}
			}
		}
		if t == true {
			break
		}
	}
	t = false
	fmt.Print("индекс   ")
	for i := range e {
		fmt.Print(" ", i, " ")
	}
	fmt.Println()
	fmt.Print("энергия  ")
	for i := range e {
		fmt.Print(" ", e[i], " ")
	}
	fmt.Println()
	for i := 1; i <= n-1; i++ {
		progres = progres + i
	}
	maxsvz = (n * n) - n - progres
	fmt.Println("максимум связей может быть ", maxsvz)
	for {
		fmt.Println("введите количество связей")
		fmt.Scan(&kolsvz)
		kolsvz, t = proverkaa(kolsvz, t)
		if kolsvz > maxsvz {
			fmt.Println("количесвто связей больше нормы")
			t = false
		}
		if t == true {
			break
		}
	}
	t = false
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
				if svz[row][col] == 0 || svz[row][col] < 0 || svz[row][col] > n {
					fmt.Println("ошибка, число должно быть в диапозоне от 1 до количества жителей")
					col = col - 1
				}
			}
			if col == 1 {
				fmt.Println("введите его знакомого")
				fmt.Scan(&svz[row][col])
				if svz[row][col] == 0 || svz[row][col] < 0 || svz[row][col] > n {
					fmt.Println("ошибка, число должно быть в диапозоне от 1 до количества жителей")
					col = col - 1
				}

			}
			if col > 0 {
				if svz[row][0] == svz[row][1] {
					fmt.Println("ошибка, числа равны")
					col = col - 1
				}
			}
		}
		aa = svz[row][0]
		bb = svz[row][1]
		t = false
		kk = 0
		for roww := 0; roww < kolsvz; roww++ {
			for coll := 0; coll < 2; coll++ {
				if aa == svz[roww][1] && bb == svz[roww][0] {
					fmt.Println("ошибка, пары зеркальны")
					svz[row][1] = -1
					t = true
					break
				}
				if aa == svz[roww][0] && bb == svz[roww][1] {
					kk = kk + 1
				}
				if kk > 2 {
					fmt.Println("ошибка, пары одинаковые")
					svz[row][1] = 0
					t = true
					break
				}

			}
			if t == true {
				break
			}
		}
		if t == true {
			row = row - 1

		}
	}
	t = false
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
					if m == n*n {
						break
					}
					a[m] = svz[row][0]
					m = m + 1
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
func proverkaa(u int, r bool) (int, bool) {
	if u == 0 || u < 0 {
		fmt.Println("число либо равно 0, либо отрицательное.ошибка")
		r = false
	} else {
		r = true
	}
	return u, r
}
