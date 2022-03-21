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

func Any[T any](list []T, f func(T) bool) bool {
	for _, value := range list {
		if f(value) {
			return true
		}
	}
	return false
}

func All[T any](list []T, f func(T) bool) bool {
	for _, value := range list {
		if !f(value) {
			return false
		}
	}
	return true
}

func None[T any](list []T, f func(T) bool) bool {
	return !Any(list, f)
}

func NotAll[T any](list []T, f func(T) bool) bool {
	return !All(list, f)
}

func Reduce[T any](list []T, f func(a, b T) T) T {
	var result T
	if len(list) == 0 {
		return result
	}
	result = list[0]
	if len(list) == 1 {
		return result
	}
	for _, value := range list[1:] {
		result = f(result, value)
	}
	return result
}

func First[T any](list []T, f func(T) bool) (T, bool) {
	for _, value := range list {
		if f(value) {
			return value, true
		}
	}
	var zero T
	return zero, false
}

func Head[T any](list []T, n int) []T {
	result := []T{}
	if n < 1 {
		return result
	}
	for i, value := range list {
		if i == n {
			break
		}
		result = append(result, value)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Tail[T any](list []T, n int) []T {

	result := []T{}
	if n < 1 {
		return result
	}

	start := len(list) - n
	if start < 1 {
		return list
	}
	return list[start:]
}
