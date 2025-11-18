package problem

// Write a function that returns the first element of any slice:
func First[T any](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}
	return items[0], true
}

// Write a function that finds the bigger of two numbers.
// Use a constraint so it only works for numeric types.
type Number interface {
	~int | ~float64
}

func Max[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// Write a generic "Pair" type.
// Create a struct that stores two values of possibly different types.
type Pair[A, B any] struct {
	First  A
	Second B
}

func main() {
	_ = Pair[int, string]{
		First:  19,
		Second: "NS",
	}
}
