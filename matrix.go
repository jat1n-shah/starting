package main

import "fmt"

func main() {

	var a [2][2]int
	var b [2][2]int
	var c [2][2]int
	fmt.Println("Enter the values for matrix a")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	fmt.Println("Enter the values for matrix b")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&b[i][j])
		}
	}

	fmt.Println("Addition of two matrix = ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
			fmt.Println(c)
		}
	}

	fmt.Println("Subtraction of two matrix = ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] - b[i][j]
			fmt.Println(c)
		}
	}

}
