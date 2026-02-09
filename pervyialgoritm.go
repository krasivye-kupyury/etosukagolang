package main

import "fmt"

var k, c, t int

func main() {
	fmt.Println("введите кол-во произнесенных чисел")
	k := vvodsproverkoy()
	t = 0
	a := make([]int, k)
	fmt.Println("введите последовательность поддтягиваний")
	for index := range a {
		a[index] = vvodsproverkoy()
	}
	for index := range a {

		if a[index] == 1 {
			c = c + 1
		}
	}
	fmt.Println("количесвто подходов")
	fmt.Println(c)
	b := make([]int, c)
	fmt.Println("количесвто подтягиваний в каждом подходе")
	for j := range b {
		for i := t; i < len(a)-1; i++ {
			if a[i+1] == 1 {

				b[j] = a[i]
				t = t + 1
				break
			}
			if a[i+1] != 1 {
				t = t + 1
			}
		}
	}
	b[c-1] = a[k-1]
	for j := range b {
		fmt.Println(b[j])
	}

}
func vvodsproverkoy() int {
	var value int
	fmt.Scan(&value)
	for {
		if value <= 0 {
			fmt.Println("число отрицательное или равно 0,попробуйте ввести заново")
			fmt.Scan(&value)
		}
		if value > 0 {
			break
		}

	}
	return value
}
