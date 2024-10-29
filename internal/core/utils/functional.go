package utils

func Reduce[T any](items []T, initial float32, reducer func(float32, T) float32) float32 {
	accumulator := initial
	for _, item := range items {
		accumulator = reducer(accumulator, item)
	}
	return accumulator
}