package main

import "fmt"

func main() {
	fmt.Println("Sleces...")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(s)
	s = remove(s, 0)
	fmt.Println(s)
	s = remove(s, 9)
	fmt.Println(s)
	s = remove(s, 3)
	fmt.Println(s)
	s = remove(s, 5)
	fmt.Println(s)
}

func remove(slice []int, s int) []int {
	if s < 0 || s > len(slice) {
		// ничего удалить нельзя
		return slice
	}
	if s == 0 {
		// первый элемент
		return slice[1:]
	}
	if s == len(slice) {
		// последний элемент
		return slice[:len(slice)-1]
	}
	// средние элементы
	return append(slice[:s], slice[s+1:]...)
}
