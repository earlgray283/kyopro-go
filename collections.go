package kyopro

type Set[K comparable] map[K]struct{}

func MakeSetFromSlice[K comparable](keys []K) Set[K] {
	set := Set[K]{}
	for _, k := range keys {
		set.Insert(k)
	}
	return set
}

func (set *Set[K]) Insert(k K) {
	(*set)[k] = struct{}{}
}

func (set *Set[K]) Remove(k K) {
	delete(*set, k)
}

func (set *Set[K]) ContainsKey(k K) bool {
	_, ok := (*set)[k]
	return ok
}

func (set *Set[K]) Iter() []K {
	keys := make([]K, len(*set))
	i := 0
	for k := range *set {
		keys[i] = k
		i++
	}
	return keys
}
