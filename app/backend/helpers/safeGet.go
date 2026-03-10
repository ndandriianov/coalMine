package helpers

func SafeGet[T any](id int, array []T) (T, bool) {
	var zero T
	if id >= 0 && id < len(array) {
		return array[id], true
	}
	return zero, false
}
