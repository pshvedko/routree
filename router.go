package routree

import (
	"sort"
)

type Node[T any] struct {
	n Nodes[T]
	v []T
	u Digit
}

type Nodes[T any] []*Node[T]

func (nn Nodes[T]) Len() int {
	return len(nn)
}

func (nn Nodes[T]) Less(i, j int) bool {
	return nn[i].u < nn[j].u
}

func (nn Nodes[T]) Swap(i, j int) {
	n := nn[i]
	nn[i] = nn[j]
	nn[j] = n
}

func (nn *Nodes[T]) Add(p Pattern, v T) int {
	if len(p) == 0 {
		return 0
	}
	u, p := p[0], p[1:]
	n := nn.Get(u)
	if n == nil {
		n = &Node[T]{
			u: u,
		}
		if u&0x8000 == 0x8000 {
			n.n = append(n.n, n) // TODO unlink cyclic reference before delete
		}
		*nn = append(*nn, n)
		sort.Sort(nn)
	}
	switch len(p) {
	case 0:
		n.v = append(n.v, v)
	default:
		return n.n.Add(p, v)
	}
	return len(n.v)
}

func (nn Nodes[T]) Get(u Digit) *Node[T] {
	i := sort.Search(len(nn), func(i int) bool { return nn[i].u >= u })
	if i < len(nn) && nn[i].u == u {
		return nn[i]
	}
	return nil
}

func (nn Nodes[T]) At(i int) *Node[T] {
	return nn[i]
}

func (nn Nodes[T]) Match(p Pattern) []T {
	var vv []T
	if len(p) > 0 {
		u := p[0]
		p = p[1:]
		for _, n := range nn {
			if n.u&u&0x7FFF == u && n.u&0x4000 == u&0x4000 {
				if len(p) == 0 {
					vv = append(vv, n.v...)
				}
				vv = append(vv, n.n.Match(p)...)
			}
		}
	}
	return vv
}

func (nn Nodes[T]) MatchFunc(p Pattern, f func(T) bool) bool {
	if len(p) > 0 {
		u := p[0]
		p = p[1:]
		for _, n := range nn {
			if n.u&u&0x7FFF == u && n.u&0x4000 == u&0x4000 {
				if len(p) == 0 {
					for _, v := range n.v {
						if !f(v) {
							return false
						}
					}
				}
				if !n.n.MatchFunc(p, f) {
					return false
				}
			}
		}
	}
	return true
}

type Router[T any] struct {
	n Nodes[T]
}

func (r *Router[T]) Add(patterns []Pattern, value T) int {
	var n int
	for _, pattern := range patterns {
		n += r.n.Add(pattern, value)
	}
	return n
}

func (r Router[T]) Match(phone Pattern) []T {
	return r.n.Match(phone)
}

func (r Router[T]) MatchFunc(phone Pattern, f func(T) bool) bool {
	return r.n.MatchFunc(phone, f)
}

func (r Router[T]) Dump(f func(u Digit, v []T, l int, e bool)) {
	r.n.dump(f, 0)
}

func (nn Nodes[T]) dump(f func(Digit, []T, int, bool), l int) {
	z := len(nn) - 1
	for i, n := range nn {
		f(n.u, n.v, l, i == z)
		if n.u&0x8000 != 0x8000 {
			n.n.dump(f, l+1)
		}
	}
}

const name2 = 1 << 10
const name1 = 0x8000
