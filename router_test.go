package routree

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func Test_nodes_Add(t *testing.T) {
	type args struct {
		p Pattern
		v any
	}
	tests := []struct {
		name string
		nn   Nodes[any]
		args args
		want Nodes[any]
	}{
		// TODO: Add test cases.
		{
			nn: nil,
			args: args{
				p: []Digit{1},
				v: 1,
			},
			want: Nodes[any]{{
				n: nil,
				u: 1,
				v: []any{1},
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
		nn   Nodes[any]
		args args
		want *Node[any]
	}{
		// TODO: Add test cases.
		{
			nn: Nodes[any]{{
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
			want: &Node[any]{
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
		u Digit
	}
	tests := []struct {
		name string
		nn   Nodes[any]
		args args
		want *Node[any]
	}{
		// TODO: Add test cases.
		{
			nn: Nodes[any]{{
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
			want: &Node[any]{
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
		nn   Nodes[any]
		want int
	}{
		// TODO: Add test cases.
		{
			nn: Nodes[any]{{
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
		nn   Nodes[any]
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			nn: Nodes[any]{{
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
		nn   Nodes[any]
		args args
		want Nodes[any]
	}{
		// TODO: Add test cases.
		{
			nn: Nodes[any]{{
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
			want: Nodes[any]{{
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
	r := Router[string]{}
	for i, pattern := range map[int]string{
		0: ".*",
		1: "7495123.*",
		2: "7(49[5|9]).......*",
		3: "7(49[5|9])......*",
		4: "7(49[5|9]).......",
		5: "1(72[0-3|4-7|8|9],5[5|7].)......*",
		6: "7495#123.*",
	} {
		patterns, err := ParsePattern(pattern)
		if err != nil {
			return
		}
		r.Add(patterns, fmt.Sprintf("%d:%q", i, pattern))
	}
	for _, number := range []string{
		"74951234567",
		"74981234567",
		"74991234567",
		"749512345678",
		"7495#1234567",
		"17211234567",
		"15555555555",
	} {
		phone, err := ParsePhone(number)
		if err != nil {
			return
		}
		fmt.Printf("%-12s -> %v\n", number, r.Match(phone))
	}
	m := make(map[int]bool)
	fmt.Println("┬")
	r.Dump(func(u Digit, v []string, l int, e bool) {
		m[l] = e
		var p strings.Builder
		for i := 0; i < l; i++ {
			if m[i] {
				p.WriteString("    ")
			} else {
				p.WriteString("│   ")
			}
		}
		if e {
			p.WriteString("└──")
		} else {
			p.WriteString("├──")
		}
		fmt.Printf("%s [%s]", p.String(), bytes.Replace(u.Split(), []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, []byte{'.'}, 1))
		if len(v) != 0 {
			fmt.Printf(" = %v", v)
		}
		fmt.Println()
	})
	// Output:
	// 74951234567  -> [1:"7495123.*" 4:"7(49[5|9])......." 2:"7(49[5|9]).......*" 3:"7(49[5|9])......*" 0:".*"]
	// 74981234567  -> [0:".*"]
	// 74991234567  -> [4:"7(49[5|9])......." 2:"7(49[5|9]).......*" 3:"7(49[5|9])......*" 0:".*"]
	// 749512345678 -> [1:"7495123.*" 2:"7(49[5|9]).......*" 3:"7(49[5|9])......*" 0:".*"]
	// 7495#1234567 -> [6:"7495#123.*"]
	// 17211234567  -> [5:"1(72[0-3|4-7|8|9],5[5|7].)......*" 0:".*"]
	// 15555555555  -> [5:"1(72[0-3|4-7|8|9],5[5|7].)......*" 0:".*"]
	// ┬
	// ├── [1]
	// │   ├── [5]
	// │   │   └── [57]
	// │   │       └── [.]
	// │   │           └── [.]
	// │   │               └── [.]
	// │   │                   └── [.]
	// │   │                       └── [.]
	// │   │                           └── [.]
	// │   │                               └── [.*] = [5:"1(72[0-3|4-7|8|9],5[5|7].)......*"]
	// │   └── [7]
	// │       └── [2]
	// │           └── [.]
	// │               └── [.]
	// │                   └── [.]
	// │                       └── [.]
	// │                           └── [.]
	// │                               └── [.]
	// │                                   └── [.*] = [5:"1(72[0-3|4-7|8|9],5[5|7].)......*"]
	// ├── [7]
	// │   └── [4]
	// │       └── [9]
	// │           ├── [5]
	// │           │   └── [1]
	// │           │       └── [2]
	// │           │           └── [3]
	// │           │               └── [.*] = [1:"7495123.*"]
	// │           ├── [59]
	// │           │   └── [.]
	// │           │       └── [.]
	// │           │           └── [.]
	// │           │               └── [.]
	// │           │                   └── [.]
	// │           │                       ├── [.]
	// │           │                       │   ├── [.] = [4:"7(49[5|9])......."]
	// │           │                       │   └── [.*] = [2:"7(49[5|9]).......*"]
	// │           │                       └── [.*] = [3:"7(49[5|9])......*"]
	// │           └── [5#]
	// │               └── [1]
	// │                   └── [2]
	// │                       └── [3]
	// │                           └── [.*] = [6:"7495#123.*"]
	// └── [.*] = [0:".*"]
}

func makeRouter() (r Router[int]) {
	for u0 := 0; u0 < 10; u0++ {
		for u1 := 0; u1 < 10; u1++ {
			for u2 := 0; u2 < 10; u2++ {
				for u3 := 0; u3 < 10; u3++ {
					r.Add([]Pattern{{1 << u0, 1 << u1, 1 << u2, 1 << u3, 0x3FF | 0x8000}}, u0*u1*u2*u3)
				}
			}
		}
	}
	return
}

func BenchmarkRouter_Add(b *testing.B) {
	r := Router[int]{}
	for i := 0; i < b.N; i++ {
		u0 := i / 1 % 10
		u1 := i / 10 % 10
		u2 := i / 100 % 10
		u3 := i / 1000 % 10
		r.Add([]Pattern{{1 << u0, 1 << u1, 1 << u2, 1 << u3, 0x3FF | 0x8000}}, u0*u1*u2*u3)
	}
}

func BenchmarkRouter_Match(b *testing.B) {
	r := makeRouter()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u3 := i / 1 % 10
		u2 := i / 10 % 10
		u1 := i / 100 % 10
		u0 := i / 1000 % 10
		v := r.Match(Pattern{1 << u0, 1 << u1, 1 << u2, 1 << u3, 1, 2, 4, 8, 16, 32})
		if len(v) != 1 {
			b.Fatalf("result length %d", len(v))
		}
		if v[0] != u0*u1*u2*u3 {
			b.Fatalf("result %d != %d", v[0], u0*u1*u2*u3)
		}
	}
}

func BenchmarkRouter_MatchFunc(b *testing.B) {
	r := makeRouter()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u3 := i / 1 % 10
		u2 := i / 10 % 10
		u1 := i / 100 % 10
		u0 := i / 1000 % 10
		var count int
		var found int
		r.MatchFunc(Pattern{1 << u0, 1 << u1, 1 << u2, 1 << u3, 1, 2, 4, 8, 16, 32}, func(value int) bool {
			found = value
			count++
			return false
		})
		if count != 1 {
			b.Fatalf("result length %d", count)
		}
		if found != u0*u1*u2*u3 {
			b.Fatalf("result %d != %d", found, u0*u1*u2*u3)
		}
	}
}

func genPhone(length int) string {
	digits := "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	return string(b)
}

func BenchmarkRouter_Match_Random(b *testing.B) {
	r := Router[int]{}
	count := 100_000

	for i := 0; i < count; i++ {
		pattern := fmt.Sprintf("7%s.*", genPhone(rand.Intn(7)+3))
		patterns, _ := ParsePattern(pattern)
		r.Add(patterns, i)
	}

	testPhones := make([]Pattern, 1000)
	for i := 0; i < 1000; i++ {
		p, _ := ParsePhone("7" + genPhone(10))
		testPhones[i] = p
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Match(testPhones[i%1000])
	}
}

func TestRouter_Longest_Prefix_Match(t *testing.T) {
	type args struct {
		r Router[int]
		p []Pattern
		q Pattern
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				r: Router[int]{},
				p: []Pattern{
					0: {1, 0x3FF, 4}, // <- last
					1: {1, 2, 0x3FF}, //
					2: {1, 2, 4},     // <- most preferred
					3: {1, 2, 6},     // <- next match
				},
				q: Pattern{1, 2, 4},
			},
			want:  []int{2, 3, 1, 0},
			want1: []int{1, 1, 1, 1},
		},
		{
			name: "",
			args: args{
				r: Router[int]{},
				p: []Pattern{
					0: {0x3FF | 0x8000},
					1: {0x3FF, 0x3FF, 0x3FF | 0x8000}, // <- next match
					2: {0x3FF, 0x3FF, 0x3FF},          // <- most preferred
					3: {0x3FF, 0x3FF, 0x3FF},          // <- most preferred, second value
				},
				q: Pattern{1, 2, 4},
			},
			want:  []int{2, 3, 1, 0},
			want1: []int{1, 1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for v := range tt.args.p {
				if got1 := tt.args.r.Add(tt.args.p[v:v+1], v); got1 != tt.want1[v] {
					t.Errorf("Add() = %v, want %v", got1, tt.want1[v])
				}
			}
			if got := tt.args.r.Match(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
