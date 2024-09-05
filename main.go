package main

import "fmt"

func main() {
	fmt.Println("Sleces...")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(s)
	s = remove(s, 5)
	fmt.Println(s)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func remove1(slice []int, s int) []int {
	result := make([]int, 0)
	result = append(slice[:s], slice[s+1:]...)
	return result
}
