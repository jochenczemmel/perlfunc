package lists

type List[T any] struct {
	Values []T
}

func New[T any](values ...T) *List[T] {
	return &List[T]{
		Values: values,
	}
}

func (l *List[T]) Len() int {
	return len(l.Values)
}

func (l *List[T]) Push(value ...T) {
	Push(&l.Values, value...)
}

func (l *List[T]) Pop() T {
	return Pop(&l.Values)
}

func (l *List[T]) Unshift(value ...T) {
	Unshift(&l.Values, value...)
}

func (l *List[T]) Shift() T {
	return Shift(&l.Values)
}

func (l List[T]) Grep(f func(T) bool) *List[T] {
	return &List[T]{
		Values: Grep(l.Values, f),
	}
}

func (l List[T]) Map(f func(T) T) *List[T] {
	return &List[T]{
		Values: Map(l.Values, f),
	}
}

func (l List[T]) All(f func(T) bool) bool {
	return All(l.Values, f)
}

func (l List[T]) Any(f func(T) bool) bool {
	return Any(l.Values, f)
}

func (l List[T]) NotAll(f func(T) bool) bool {
	return NotAll(l.Values, f)
}

func (l List[T]) None(f func(T) bool) bool {
	return None(l.Values, f)
}

func (l List[T]) Reduce(f func(a, b T) T) T {
	return Reduce(l.Values, f)
}
