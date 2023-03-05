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
	if len(p) == 0 {
		return
	}
	n := nn.Get(p[0])
	if n == nil {
		n = new(Node)
		n.u = p[0]
		*nn = append(*nn, n)
		sort.Sort(nn)
	}
	if len(p) == 1 {
		n.v = v
		return
	}
	n.n.Add(p[1:], v)
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
				switch {
				case len(p) == 1 && n.v != nil:
					return []interface{}{n.v}
				case n.n != nil:
					vv = append(vv, n.n.Match(p[1:])...)
				}
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
