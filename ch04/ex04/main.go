package main

import "fmt"

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	n := 4

	l := len(arr)
	g := gcd(l, n)

	fmt.Printf("before: %v\n", arr)
	for i := 0; i < g; i++ {
		for j := (i + n) % l; j != i; j = (j + n) % l {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	fmt.Printf("after: %v\n", arr)
}
