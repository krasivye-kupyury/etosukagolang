package main

import "fmt"

func main() {
	var k, c, t int
	fmt.Println("введите кол-во произнесенных чисел")
	fmt.Scan(&k)
	t = 0
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
