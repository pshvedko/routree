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
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nn.Add(tt.args.p, tt.args.v)
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
		{},
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
		{},
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
		{},
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
		{},
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
		// FIXME
		".*",
		"7(49[5|9]).......",
		"1(72[0-9]).......",
	} {
		p, err := ParseString(pattern)
		if err != nil {
			return
		}
		r.Add(p, map[string]interface{}{"pattern": pattern})
	}
	phone1, err := ParseString("74951234567")
	if err != nil {
		return
	}
	phone2, err := ParseString("17211234567")
	if err != nil {
		return
	}
	fmt.Println(r.Match(phone1[0]))
	fmt.Println(r.Match(phone2[0]))
	// Output:
	// [map[pattern:7(49[5|9]).......]]
	// [map[pattern:1(72[0-9]).......]]
}
