package routree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_nodes_Add(t *testing.T) {
	type args struct {
		p Pattern
		v interface{}
	}
	tests := []struct {
		name string
		nn   Nodes
		args args
		want Nodes
	}{
		// TODO: Add test cases.
		{
			nn: nil,
			args: args{
				p: []uint16{1},
				v: nil,
			},
			want: Nodes{{
				n: nil,
				u: 1,
				v: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nn.Add(tt.args.p, tt.args.v)
			if !reflect.DeepEqual(tt.nn, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.nn, tt.want)
			}
		})
	}
}

func Test_nodes_At(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		nn   Nodes
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			nn: Nodes{{
				n: nil,
				u: 0,
				v: nil,
			}, {
				n: nil,
				u: 1,
				v: nil,
			}, {
				n: nil,
				u: 2,
				v: nil,
			}},
			args: args{
				i: 1,
			},
			want: &Node{
				n: nil,
				u: 1,
				v: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nn.At(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodes_Get(t *testing.T) {
	type args struct {
		u uint16
	}
	tests := []struct {
		name string
		nn   Nodes
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			nn: Nodes{{
				n: nil,
				u: 0,
				v: nil,
			}, {
				n: nil,
				u: 1,
				v: nil,
			}, {
				n: nil,
				u: 2,
				v: nil,
			}},
			args: args{
				u: 1,
			},
			want: &Node{
				n: nil,
				u: 1,
				v: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nn.Get(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodes_Len(t *testing.T) {
	tests := []struct {
		name string
		nn   Nodes
		want int
	}{
		// TODO: Add test cases.
		{
			nn: Nodes{{
				n: nil,
				u: 0,
				v: nil,
			}, {
				n: nil,
				u: 1,
				v: nil,
			}, {
				n: nil,
				u: 2,
				v: nil,
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nn.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodes_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		nn   Nodes
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			nn: Nodes{{
				n: nil,
				u: 0,
				v: nil,
			}, {
				n: nil,
				u: 1,
				v: nil,
			}, {
				n: nil,
				u: 2,
				v: nil,
			}},
			args: args{
				i: 1,
				j: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nn.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodes_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		nn   Nodes
		args args
		want Nodes
	}{
		// TODO: Add test cases.
		{
			nn: Nodes{{
				n: nil,
				u: 0,
				v: nil,
			}, {
				n: nil,
				u: 1,
				v: nil,
			}, {
				n: nil,
				u: 3,
				v: nil,
			}},
			args: args{
				i: 0,
				j: 1,
			},
			want: Nodes{{
				n: nil,
				u: 1,
				v: nil,
			}, {
				n: nil,
				u: 0,
				v: nil,
			}, {
				n: nil,
				u: 3,
				v: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nn.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.nn, tt.want) {
				t.Errorf("Swap() = %v, want %v", tt.nn, tt.want)
			}
		})
	}
}

func ExampleRouter_Add() {
	r := Router{}
	for _, pattern := range []string{
		".*",
		"7495123.*",
		"7(49[5|9]).......*",
		"7(49[5|9])......*",
		"7(49[5|9]).......",
		"1(72[0-9])......*",
	} {
		p, err := ParseString(pattern)
		if err != nil {
			return
		}
		r.Add(p, fmt.Sprintf("/%s/:%d", pattern, len(pattern)))
	}
	for _, number := range []string{
		"74951234567",
		"74981234567",
		"74991234567",
		"749512345678",
		"17211234567",
	} {
		phone, err := ParseString(number)
		if err != nil {
			return
		}
		fmt.Printf("%-16s -> %v\n", number, r.Match(phone[0]))
	}
	// Output:
	// 74951234567      -> [/7495123.*/:9 /7(49[5|9])......./:17 /7(49[5|9]).......*/:18 /7(49[5|9])......*/:17 /.*/:2]
	// 74981234567      -> [/.*/:2]
	// 74991234567      -> [/7(49[5|9])......./:17 /7(49[5|9]).......*/:18 /7(49[5|9])......*/:17 /.*/:2]
	// 749512345678     -> [/7495123.*/:9 /7(49[5|9]).......*/:18 /7(49[5|9])......*/:17 /.*/:2]
	// 17211234567      -> [/1(72[0-9])......*/:17 /.*/:2]
}
