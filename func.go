package main

import "fmt"

func circle(r float32) float32 {
	return r * 3.14
}

func square(a int16) int16 {
	return a * a
}

func rectangle(a, b int16) int16 {
	return a * b
}

func main() {

	var a, b int16
	var r float32

	fmt.Println("enter the radius")
	fmt.Scan(&r)

	fmt.Println("enter the length")
	fmt.Scan(&a)

	fmt.Println("enter the width")
	fmt.Scan(&b)

	s := circle(r)
	fmt.Println("area of circle=", s)

	t := square(a)
	fmt.Println("area of square =", t)

	u := rectangle(a, b)
	fmt.Println("area of rectangle=", u)

}
