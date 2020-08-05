package set

type StringSet map[string]bool

func (s StringSet) Contains(item string) bool {
	_, ok := s[item]
	return ok
}

func FromSlice(items []string) StringSet {
	res := make(StringSet, len(items))
	for _, item := range items {
		res[item] = true
	}
	return res
}
