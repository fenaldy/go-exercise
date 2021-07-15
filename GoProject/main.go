package main

import (
	htg "GoProject/hitung"
	"fmt"
)

func Add(numa int, numb int) int {
	return numa + numb
}

func main() {

	var angka = []int{1, 2, 3}
	for a, b := range angka {
		fmt.Println(a, "-->", b)
	}
	fmt.Println("test")

	x := htg.Tambah(4, 4)
	fmt.Println(x)
}
