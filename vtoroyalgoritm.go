package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var letN, letW, letS, letE = "n", "w", "s", "e"
var min = 0.001
var max = 99.999

func main() {
	rand.Seed(time.Now().UnixNano())
	var s string
	var valueN, valueW, valueS, valueE, x, y = 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	var countN, countW, countS, countE = 0, 0, 0, 0
	var sumN, sumW, sumS, sumE float64
	var n, w, ss, e, b = false, false, false, false, false
	for {
		fmt.Println("введите строку")
		fmt.Scan(&s)
		s, b = proverka(s, b)
		if b == true {
			break
		}
	}
	countN, n = schetbukv(letN, s)
	countE, e = schetbukv(letE, s)
	countS, ss = schetbukv(letS, s)
	countW, w = schetbukv(letW, s)
	if e == true {
		if w == true {
			valueE, valueW, sumE, sumW = schetkoord1(countE, countW)
			w = false
		} else {
			valueE, sumE = schetkoord2(countE)
		}
	}
	if n == true {
		if ss == true {
			valueN, valueS, sumN, sumS = schetkoord1(countN, countS)
			ss = false
		} else {
			valueN, sumN = schetkoord2(countN)
		}
	}
	if w == true {
		if e == true {
			valueW, valueE, sumW, sumW = schetkoord1(countW, countE)
			e = false
		} else {
			valueW, sumW = schetkoord2(countW)
		}
	}
	if ss == true {
		if n == true {
			valueS, valueN, sumS, sumN = schetkoord1(countS, countN)
			n = false
		} else {
			valueS, sumS = schetkoord2(countS)
		}
	}
	fmt.Println("N =", valueN, "kolich N =", countN, "summa N=", sumN)
	fmt.Println("W =", valueW, "kolich W =", countW, "summa W=", sumW)
	fmt.Println("S =", valueS, "kolich S =", countS, "summa S=", sumS)
	fmt.Println("E =", valueE, "kolich E =", countE, "summa E=", sumE)
	for _, runevalue := range s {
		if string(runevalue) == letN {
			y = y + valueN
			fmt.Printf("%.6f %.6f\n", x, y)
		}
		if string(runevalue) == letW {
			x = x - valueW
			fmt.Printf("%.6f %.6f\n", x, y)
		}
		if string(runevalue) == letS {
			y = y - valueS
			fmt.Printf("%.6f %.6f\n", x, y)
		}
		if string(runevalue) == letE {
			x = x + valueE
			fmt.Printf("%.6f %.6f\n", x, y)
		}
	}
	if x > 0 {
		x = math.Floor(x)
	}
	if y > 0 {
		y = math.Floor(y)
	}
	if x < 0 {
		x = math.Ceil(x)
	}
	if y < 0 {
		y = math.Ceil(y)
	}
	if x == 0 && y == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
func proverka(stroka string, bb bool) (string, bool) {
	for _, runevalue := range stroka {
		if (string(runevalue) != letN) && (string(runevalue) != letW) && (string(runevalue) != letS) && (string(runevalue) != letE) {
			fmt.Println("ошибка, можно вводить только w s e n")
			fmt.Println("попробуйте еще раз")
			bb = false
			break
		} else {
			bb = true
		}
	}
	return stroka, bb
}
func schetbukv(bukva string, ss string) (int, bool) {
	var count = 0
	var bb bool
	for _, runevalue := range ss {
		if string(runevalue) == bukva {
			bb = true
			count = count + 1
		}
	}
	return count, bb
}
func schetkoord1(countA int, countB int) (float64, float64, float64, float64) {
	var valueA, valueB, sumA, sumB float64
	valueA = math.Round((min+rand.Float64()*(max-min))*1_000_000) / 1_000_000
	sumA = valueA * float64(countA)
	sumB = sumA
	valueB = sumB / float64(countB)
	return valueA, valueB, sumA, sumB
}
func schetkoord2(countA int) (float64, float64) {
	var valueA, sumA float64
	valueA = math.Round((min+rand.Float64()*(max-min))*1_000_000) / 1_000_000
	sumA = valueA * float64(countA)
	return valueA, sumA
}
