package gocache


type Store struct{
	items map[string]StoreValue
}

type StoreValue struct {
	name string
	value interface{}
}

func (s *Store) Make() {
	s.items = make(map[string]StoreValue)
}


func (s *Store) Set(k string, value interface{}) {
	s.items[k] = StoreValue{
		name: k,
		value: value,
	}
}

func (s *Store) Get(k string) interface{} {
	return s.items[k].value
}

func (s *Store) Count() int {
	return len(s.items) 
}

func (s *Store) Clear() {
	s.Make()
}