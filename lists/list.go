package lists

func Push[T any](list *[]T, value ...T) {
	*list = append(*list, value...)
}

func Pop[T any](list *[]T) T {
	var zero T
	if len(*list) == 0 {
		return zero
	}
	last := len(*list) - 1
	defer func() {
		*list = (*list)[:last]
	}()
	return (*list)[last]
}

func Unshift[T any](list *[]T, value ...T) {
	*list = append(value, *list...)
}

func Shift[T any](list *[]T) T {
	var zero T
	if len(*list) < 1 {
		return zero
	}
	defer func() {
		*list = (*list)[1:]
	}()
	return (*list)[0]
}

func Grep[T any](list []T, f func(T) bool) []T {
	result := []T{}
	for _, value := range list {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func Map[T any](list []T, f func(T) T) []T {
	result := []T{}
	for _, value := range list {
		result = append(result, f(value))
	}
	return result
}
