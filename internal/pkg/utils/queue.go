package utils

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) TryDequeue() (success bool, item *T) {
	if q.Len() < 1 {
		return false, nil
	}

	i := &q.items[0]
	q.items = q.items[1:]

	return true, i
}

func (q *Queue[T]) TryPeek() (success bool, item *T) {
	if q.Len() < 1 {
		return false, nil
	}
	return true, &q.items[0]
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}
