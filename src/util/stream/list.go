package stream

type IListStream[T any] <-chan T

func NewListStream[T any](items []T) IListStream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, item := range items {
			ch <- item
		}
	}()
	return ch
}

func (strm IListStream[T]) Map(fn func(T) any) IListStream[any] {
	out := make(chan any)
	go func() {
		defer close(out)
		for item := range strm {
			out <- fn(item)
		}
	}()
	return out
}

func (strm IListStream[T]) Filter(fn func(T) bool) IListStream[T] {
	out := make(chan T)
	go func() {
		defer close(out)
		for item := range strm {
			if fn(item) {
				out <- item
			}
		}
	}()
	return out
}

func (strm IListStream[T]) ToList() []T {
	out := make([]T, 0)
	for item := range strm {
		out = append(out, item)
	}
	return out
}
