package gohaskell

// Returns the first item in a  slice
func Head[T any](slice []T) T {
	return slice[0]
}

func Last[T any](xs []T) T {
	if len(xs) == 0 {
		var x T
		return x
	}
	return xs[len(xs)-1]
}

// Returns the items after the first item in a slice
func Tail[T any](slice []T) []T {
	return slice[1:]
}

// Removes an item from a slice
func Take[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func Reverse[T any](xs []T) []T {
	if len(xs) == 0 {
		return xs
	}
	return append(Reverse(Tail(xs)), Head(xs))
}

// Returns the Factorial of any number
func Factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}

// Removes any item after x
func Drop[T any](xs []T, x int) []T {
	return xs[x:]
}

// Sums up elements of an array
func Sum(xs []any) any {
	if len(xs) == 0 {
		var x any
		return x
	}
	//a := T + T
	return Head(xs).(int) + Sum(Tail(xs)).(int) //no new head
}

func contains[T comparable](xs []T, elem T) bool {
	if elem == Head(xs) {
		return true
	} else {
		if elem != Head(xs) {
			return contains(Tail(xs), elem)
		}
	}

	return false

}
