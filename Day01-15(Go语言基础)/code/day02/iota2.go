package main

import "fmt"

func main() {
	const (
		A = iota // 0
		B // 1
		C // 2
		D = "haha" // iota = 3
		E // haha iota = 4
		F = 100 //iota =5
		G //100 iota = 6
		H = iota // 7
		I //iota 8
	)
	const (
		J = iota // 0
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)
}
