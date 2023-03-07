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
	u, p := p[0], p[1:]
	n := nn.Get(u)
	if n == nil {
		n = &Node{
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
		u := p[0]
		p = p[1:]
		for _, n := range nn {
			if n.u&u&0x7FFF == u && n.u&0x4000 == u&0x4000 {
				if len(p) == 0 && n.v != nil {
					vv = append(vv, n.v)
				}
				vv = append(vv, n.n.Match(p)...)
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
