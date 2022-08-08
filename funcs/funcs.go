package main

import "fmt"

//Returns the first item in a  slice
func Head[T any](slice []T) T {
	return slice[0]
}

//Returns the items after the first item in a slice
func Tail[T any](slice []T) []T {
	return slice[1:]
}

//Removes an item from a slice
func Take[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

//func Reverse[T any](slice []T,slice2 []T) []T {
//	return Reverse()
//}

//Returns the Factorial of any number
func Factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}

//Removes any item after x
func Drop[T any](xs []T, x int) []T {
	return xs[x:]
}

//Sums up elements of an array
func Sum(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	return Head(xs) + Sum(Tail(xs)) //no new head
}

func Max[T any](xs []T) {
	return
}

func main() {
	arr := []float64{1, 2}
	ar2 := Sum(arr)
	fmt.Println(ar2)
}
