package main

func Sum[T int](num []T) T {
	total := T(0)
	for _, v := range num {
		total += v
	}

	return total
}
