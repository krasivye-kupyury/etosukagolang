package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0.001
	max := 99.999
	var s string
	var n, w, ss, e, b bool
	b = false
	n = false
	w = false
	ss = false
	e = false
	var valueN, valueW, valueS, valueE, x, y float64
	var countN, countW, countS, countE int64
	var sumN, sumW, sumS, sumE float64
	x = 0.0
	y = 0.0
	countN = 0
	countW = 0
	countS = 0
	countE = 0
	valueE = 0
	valueN = 0
	valueS = 0
	valueW = 0
	for {
		if b == false {
			fmt.Println("введите строку")
			fmt.Scan(&s)
			for _, runevalue := range s {
				if (string(runevalue) != "n") && (string(runevalue) != "w") && (string(runevalue) != "s") && (string(runevalue) != "e") {
					fmt.Println("ошибка, можно вводить только w s e n")
					b = false
					break
				} else {
					b = true
				}
			}
		} else {
			break
		}
	}
	for _, runevalue := range s {
		if string(runevalue) == "n" {
			n = true
			countN = countN + 1
		}
		if string(runevalue) == "w" {
			w = true
			countW = countW + 1
		}
		if string(runevalue) == "s" {
			ss = true
			countS = countS + 1
		}
		if string(runevalue) == "e" {
			e = true
			countE = countE + 1
		}
	}
	if e == true {
		if w == true {
			valueE = min + rand.Float64()*(max-min)
			sumE = valueE * float64(countE)
			sumW = sumE
			valueW = sumW / float64(countW)
			w = false
		} else {
			valueE = min + rand.Float64()*(max-min)
		}
	}
	if n == true {
		if ss == true {
			valueN = min + rand.Float64()*(max-min)
			sumN = valueN * float64(countN)
			sumS = sumN
			valueS = sumS / float64(countS)
			ss = false
		} else {
			valueN = min + rand.Float64()*(max-min)
		}
	}
	if w == true {
		if e == true {
			valueW = min + rand.Float64()*(max-min)
			sumW = valueW * float64(countW)
			sumE = sumW
			valueE = sumE / float64(countE)
			e = false
		} else {
			valueW = min + rand.Float64()*(max-min)
		}
	}
	if ss == true {
		if n == true {
			valueS = min + rand.Float64()*(max-min)
			sumS = valueS * float64(countS)
			sumN = sumS
			valueN = sumN / float64(countN)
			n = false
		} else {
			valueS = min + rand.Float64()*(max-min)
		}
	}
	fmt.Println("N =", valueN, "kolich N =", countN, "summa N=", sumN)
	fmt.Println("W =", valueW, "kolich W =", countW, "summa W=", sumW)
	fmt.Println("S =", valueS, "kolich S =", countS, "summa S=", sumS)
	fmt.Println("E =", valueE, "kolich E =", countE, "summa E=", sumE)
	for _, runevalue := range s {
		if string(runevalue) == "n" {
			y = y + valueN
			fmt.Println(x, " ; ", y)
		}
		if string(runevalue) == "w" {
			x = x - valueW
			fmt.Println(x, " ; ", y)
		}
		if string(runevalue) == "s" {
			y = y - valueS
			fmt.Println(x, " ; ", y)
		}
		if string(runevalue) == "e" {
			x = x + valueE
			fmt.Println(x, " ; ", y)
		}
	}
	if x == 0 && y == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
		fmt.Println("если у одной из координат на конце ...е-..")
		fmt.Println("то это обозначает 0 и на самом деле вывод должен быть yes")
		fmt.Println("это происходит из-за погрешности при работе с float, как это пофиксить я пока хз")
	}
}
