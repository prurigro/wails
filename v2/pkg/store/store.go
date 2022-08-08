package store

type storable interface {
}

type Store[T any] struct {
	id          string
	data        T
	subscribers []func(T)
}

func New[T any](id string, data T) *Store[T] {
	return &Store[T]{
		id:   id,
		data: data,
	}
}

func (s *Store[T]) notifySubscribers() {
	for _, subscriber := range s.subscribers {
		subscriber(s.data)
	}
}

func (s *Store[T]) Get() T {
	return s.data
}

func (s *Store[T]) Set(data T) {
	s.data = data
	s.notifySubscribers()
}

func (s *Store[T]) Update(f func(T) T) {
	s.Set(f(s.data))
}

func (s *Store[T]) Subscribe(subscriber func(T)) {
	s.subscribers = append(s.subscribers, subscriber)
}
