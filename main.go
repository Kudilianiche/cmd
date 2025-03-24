package main

import "fmt"

var h, w int

func main() {
	fmt.Println("Введите высоту и ширину шахматной доски")
	fmt.Scan(&h, &w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
