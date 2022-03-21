package perlfunc

import "github.com/jochenczemmel/perlfunc/lists"

type List[T any] struct {
	values []T
}

func NewList[T any](values ...T) *List[T] {
	return &List[T]{
		values: values,
	}
}

func (l *List[T]) Len() int {
	return len(l.values)
}

func (l *List[T]) Push(value ...T) {
	lists.Push(&l.values, value...)
	// l.values = append(l.values, value...)
}

func (l *List[T]) Pop() T {
	return lists.Pop(&l.values)
	//	var zero T
	//	if len(l.values) < 1 {
	//		return zero
	//	}
	//	last := len(l.values) - 1
	//	defer func() {
	//		l.values = l.values[:last]
	//	}()
	//	return l.values[last]
}

func (l *List[T]) Unshift(value ...T) {
	lists.Unshift(&l.values, value...)
	// l.values = append(value, l.values...)
}

func (l *List[T]) Shift() T {
	return lists.Shift(&l.values)
	//	var zero T
	//	if len(l.values) < 1 {
	//		return zero
	//	}
	//	defer func() {
	//		l.values = l.values[1:]
	//	}()
	//	return l.values[0]
}

func (l List[T]) Grep(f func(T) bool) *List[T] {
	return &List[T]{
		values: lists.Grep(l.values, f),
	}
	//	result := &List[T]{}
	//	for _, value := range l.values {
	//		if f(value) {
	//			result.values = append(result.values, value)
	//		}
	//	}
	//	return result
}

func (l List[T]) Map(f func(T) T) *List[T] {
	return &List[T]{
		values: lists.Map(l.values, f),
	}
	//	result := &List[T]{}
	//	for _, value := range l.values {
	//		result.values = append(result.values, f(value))
	//	}
	//	return result
}

func (l List[T]) Values() []T {
	result := make([]T, len(l.values))
	copy(result, l.values)
	return result
}
