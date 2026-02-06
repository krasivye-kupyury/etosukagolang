package main

import "fmt"

func main() {
	var k, c int
	fmt.Println("введите кол-во произнесенных чисел")
	fmt.Scan(&k)
	a := make([]int, k)
	fmt.Println("введите последовательность поддтягиваний")
	for index := range a {
		fmt.Scan(&a[index])
	}
	for index := range a {

		if a[index] == 1 {
			c = c + 1
		}
	}
	fmt.Println("количесвто подходов")
	fmt.Println(c)
	b := make([]int, c)
	fmt.Println("массив:")

	for i := 0; i < len(a); i++ {

		for j := range b {
			if a[i+1] == 1 {
				b[j] = a[i]
			}

		}

	}
	for index := range b {
		fmt.Println(b[index])
	}

}
