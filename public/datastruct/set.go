package datastruct

type parseKeyFunc func(e interface{}) (key interface{})

type Set struct {
	data    map[interface{}]interface{}
	keyFunc parseKeyFunc
}

func NewSetWithParseKeyFunc(f parseKeyFunc) *Set {
	set := new(Set)
	set.data = make(map[interface{}]interface{})
	set.keyFunc = f
	return set
}

func NewSet() *Set {
	set := new(Set)
	set.data = make(map[interface{}]interface{})
	return set
}

func (s Set) Exists(e interface{}) bool {
	var exists bool
	if s.keyFunc != nil {
		key := s.keyFunc(e)
		_, exists = s.data[key]
	} else {
		_, exists = s.data[e]
	}
	return exists
}

func (s *Set) add(e interface{}) {
	if s.keyFunc != nil {
		key := s.keyFunc(e)
		if _, exists := s.data[key]; !exists {
			s.data[key] = e
		}
	} else {
		s.data[e] = nil
	}
}

func (s *Set) Add(e interface{}) {
	if s.Exists(e) {
		s.Remove(e)
	}
	s.add(e)
}

func (s *Set) Remove(e interface{}) {
	if s.keyFunc != nil {
		key := s.keyFunc(e)
		delete(s.data, key)
	} else {
		delete(s.data, e)
	}
}

func (s Set) Len() int {
	return len(s.data)
}

func (s Set) ToArr() []interface{} {
	var arr []interface{}
	if s.keyFunc != nil {
		for _, v := range s.data {
			arr = append(arr, v)
		}
	} else {
		for k := range s.data {
			arr = append(arr, k)
		}
	}
	return arr
}

func (s Set) ToMap() map[interface{}]interface{} {
	return s.data
}
