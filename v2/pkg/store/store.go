package store

// Store is our main store struct
type Store[T any] struct {
	id          int
	data        T
	subscribers []func(T)
}

// New creates a new store
func New[T any](id int, data T) *Store[T] {
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

// Get returns the data in the store
func (s *Store[T]) Get() T {
	return s.data
}

// Set sets the data in the store
func (s *Store[T]) Set(data T) {
	s.data = data
	s.notifySubscribers()
}

// Update updates the data in the store with whatever is returned by given method
func (s *Store[T]) Update(f func(T) T) {
	s.Set(f(s.data))
}

// Subscribe to store updates
func (s *Store[T]) Subscribe(subscriber func(T)) {
	s.subscribers = append(s.subscribers, subscriber)
}
