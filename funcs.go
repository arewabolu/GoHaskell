package gohaskell

type arithmetic interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

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
	if len(slice) == 0 {
		return []T{}
	}
	if len(slice) == 1 {
		return slice
	}
	return slice[1:]
}

// Removes an item from a slice
func Pop[T any](slice []T, index int) []T {
	if len(slice) == 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// flips the items of slice xs
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

// Removes any item before x
func Drop[T any](xs []T, x int) []T {
	if len(xs) == 0 {
		return xs
	}
	if len(xs) < x {
		return []T{}
	}
	return xs[x:]
}

// Sums up elements of an array
func Sum[T arithmetic](xs []T) T {
	if len(xs) == 0 {
		var x T
		return x
	}
	return Head(xs) + Sum(Tail(xs))
}

// Checks if xs contains the given element elem and returns a boolean value
func Contains[T comparable](xs []T, elem T) bool {
	if len(xs) == 0 {
		return false
	}
	if elem == Head(xs) {
		return true
	}
	if elem != Head(xs) {
		return Contains(Tail(xs), elem)
	}
	return false
}

// Places item 'x' into the slice xs at position 'pos' and returns a new array containing x
func Put[T any](xs []T, x T, pos int) []T {
	if len(xs) == 0 {
		return xs
	}
	newT := make([]T, 0, len(xs)+1)
	mid := xs[:pos]
	newT = append(newT, mid...)
	newT = append(newT, x)
	newT = append(newT, xs[pos:]...)
	return newT
}

// Returns all items up to 'max' in given slice.
//
// Edge case: If 'max' is greater than the slice length it returns all items in the slice
func Take[T any](xs []T, max int) []T {
	if len(xs) == 0 {
		return xs
	}
	if len(xs) < max {
		return xs
	}
	return xs[0:max]
}

// returns all items except the last items in the slice
func ExlcudeLast[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}
