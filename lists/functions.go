package lists

// Push appends the new element(s) to the list.
func Push[T any](list *[]T, value ...T) {
	*list = append(*list, value...)
}

// Pop removes the last element from the list and returns it.
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

// Push prepends the new element(s) to the list.
func Unshift[T any](list *[]T, value ...T) {
	*list = append(value, *list...)
}

// Shift removes the first element from the list and returns it.
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

// Grep returns a list of elements that return true for f.
// f gets only the list values, not the indexes.
func Grep[L ~[]T, T any](list L, f func(T) bool) L {
	result := L{}
	for _, value := range list {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// Find returns a the first element that return true for f and true.
// If no element was found, it returns the zero value and false.
// f gets only the list values, not the indexes.
// func Find[T any](list []T, f func(T) bool) (T, bool) {
func Find[L ~[]T, T any](list L, f func(T) bool) (T, bool) {
	var zero T
	for _, value := range list {
		if f(value) {
			return value, true
		}
	}
	return zero, false
}

// Map returns the list where the values are modified by f.
// f gets only the list values, not the indexes.
func Map[L ~[]T, T any](list L, f func(T) T) L {
	result := L{}
	for _, value := range list {
		result = append(result, f(value))
	}
	return result
}

// Any returns true if f returns true for at least one element.
// f gets only the list values, not the indexes.
func Any[T any](list []T, f func(T) bool) bool {
	for _, value := range list {
		if f(value) {
			return true
		}
	}
	return false
}

// Any returns true if f returns true for all elements.
// f gets only the list values, not the indexes.
func All[T any](list []T, f func(T) bool) bool {
	for _, value := range list {
		if !f(value) {
			return false
		}
	}
	return true
}

// None returns true if f returns true for no element.
// f gets only the list values, not the indexes.
func None[T any](list []T, f func(T) bool) bool {
	return !Any(list, f)
}

// NotAll returns true if f returns true for at least one element.
// f gets only the list values, not the indexes.
func NotAll[T any](list []T, f func(T) bool) bool {
	return !All(list, f)
}

// Reduce applies f to all list elements.
//  If the list is empty, the zero value is returned.
//  If the list has one element, it is returned.
//  If the list has more elements:
//  - on the first iteration, the first two elements are
//    provided to the function
//  - on the next iteration, the result of the first funktion
//    call is provided as the first argument
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

// First returns the first occurence of an element and true, if the value was found.
// If the value was not found, it returns the zero value and false,
func First[T any](list []T, f func(T) bool) (T, bool) {
	for _, value := range list {
		if f(value) {
			return value, true
		}
	}
	var zero T
	return zero, false
}

// Head returns the first n elements of the list.
// If n > len(list), it returns all elements (len < n).
func Head[L ~[]T, T any](list L, n int) L {
	result := L{}
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

// Tail returns the last n elements of the list.
// If n > len(list), it returns all elements (len < n).
func Tail[L ~[]T, T any](list L, n int) L {

	result := L{}
	if n < 1 {
		return result
	}

	start := len(list) - n
	if start < 1 {
		return list
	}
	return list[start:]
}
