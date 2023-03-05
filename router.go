package routree

import (
	"sort"
)

type Node struct {
	n Nodes
	u uint16
	v interface{}
}

type Nodes []*Node

func (nn Nodes) Len() int {
	return len(nn)
}

func (nn Nodes) Less(i, j int) bool {
	return nn[i].u < nn[j].u
}

func (nn Nodes) Swap(i, j int) {
	n := nn[i]
	nn[i] = nn[j]
	nn[j] = n
}

func (nn *Nodes) Add(p Pattern, v interface{}) {
	p, u := nextDigit(p)
	if p == nil {
		return
	}
	n := nn.Get(u)
	if n == nil {
		n = &Node{
			u: u,
		}
		if u&0x8000 > 0 {
			n.n = append(n.n, n)
		}
		*nn = append(*nn, n)
		sort.Sort(nn)
	}
	switch len(p) {
	case 0:
		n.v = v
	default:
		n.n.Add(p, v)
	}
}

func (nn Nodes) Get(u uint16) *Node {
	i := sort.Search(len(nn), func(i int) bool { return nn[i].u >= u })
	if i < len(nn) && nn[i].u == u {
		return nn[i]
	}
	return nil
}

func (nn Nodes) At(i int) *Node {
	return nn[i]
}

func (nn Nodes) Match(p Pattern) []interface{} {
	var vv []interface{}
	if len(p) > 0 {
		for _, n := range nn {
			if n.u&p[0] > 0 {
				if len(p) == 1 && n.v != nil {
					vv = append(vv, n.v)
				}
				vv = append(vv, n.n.Match(p[1:])...)
			}
		}
	}
	return vv
}

type Router struct {
	n Nodes
}

func (r *Router) Add(patterns []Pattern, value interface{}) {
	for _, pattern := range patterns {
		r.n.Add(pattern, value)
	}
}

func (r Router) Match(phone Pattern) []interface{} {
	return r.n.Match(phone)
}