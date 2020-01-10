package set

type Set map[interface{}]struct{}

func newSet() Set {
	return make(Set)
}

func (set *Set) Add(i interface{}) bool {
	_, found := (*set)[i]
	if found {
		return false
	}

	(*set)[i] = struct{}{}
	return true
}

func (set *Set) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := (*set)[val]; !ok {
			return false
		}
	}
	return true
}

func (set *Set) Clear() {
	*set = newSet()
}

func (set *Set) Remove(i interface{}) {
	delete(*set, i)
}

func (set *Set) Size() int {
	return len(*set)
}

func NewSet(i ...interface{}) Set {
	set := newSet()
	for _, item := range i {
		set.Add(item)
	}
	return set
}
