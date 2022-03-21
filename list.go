package perlfunc

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
	l.values = append(l.values, value...)
}

func (l *List[T]) Pop() T {
	var zero T
	if len(l.values) < 1 {
		return zero
	}
	last := len(l.values) - 1
	defer func() {
		l.values = l.values[:last]
	}()
	return l.values[last]
}

func (l *List[T]) Unshift(value ...T) {
	l.values = append(value, l.values...)
}

func (l *List[T]) Shift() T {
	var zero T
	if len(l.values) < 1 {
		return zero
	}
	defer func() {
		l.values = l.values[1:]
	}()
	return l.values[0]
}

func (l List[T]) Grep(f func(T) bool) *List[T] {
	result := &List[T]{}
	for _, value := range l.values {
		if f(value) {
			result.values = append(result.values, value)
		}
	}
	return result
}

func (l List[T]) Map(f func(T) T) *List[T] {
	result := &List[T]{}
	for _, value := range l.values {
		result.values = append(result.values, f(value))
	}
	return result
}

func (l List[T]) Values() []T {
	result := make([]T, len(l.values))
	copy(result, l.values)
	return result
}
