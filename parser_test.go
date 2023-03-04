package routree

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func ExampleParse() {
	for _, s := range []string{
		"1.*",
		"1**",
		"11*",
		"*",
		"11.",
		"1..",
		"1.1",
		"[1]",
		"[]",
		"[0-1]",
		"[0-2]",
		"[0-9]",
		"[1-3]",
		"[1-1]",
		"[2-1]",
		"[-1]",
		"[2-]",
		"[1-3|7]",
		"[1-3|7-9]",
		"[1-3|7|9]",
		"[1|2]",
		"[1|2-4|0]",
		"1[2|4]567.*",
		//
		"12345,67890",
		"1(22.,33.)4(1,2)",
		"1(22.,33.)4(1)",
		"1(22.,33.)4()",
		"1(22.,33.)4(1,)",
		"1(22.,33.)4(,1)",
	} {
		pattern, err := Parse(bytes.NewBufferString(s))
		if err != nil {
			fmt.Println(s, err)
			continue
		}
		fmt.Println(s, pattern)
	}

	// Output:
	// 1.* [[2 1023 0]]
	// 1** illegal symbol '*'
	// 11* [[2 2 0]]
	// * [[0]]
	// 11. [[2 2 1023]]
	// 1.. [[2 1023 1023]]
	// 1.1 [[2 1023 2]]
	// [1] [[2]]
	// [] illegal symbol ']'
	// [0-1] [[3]]
	// [0-2] [[7]]
	// [0-9] [[1023]]
	// [1-3] [[14]]
	// [1-1] illegal range '1-1'
	// [2-1] illegal range '2-1'
	// [-1] illegal symbol '-'
	// [2-] illegal symbol ']'
	// [1-3|7] [[142]]
	// [1-3|7-9] [[910]]
	// [1-3|7|9] [[654]]
	// [1|2] [[6]]
	// [1|2-4|0] [[31]]
	// 1[2|4]567.* [[2 20 32 64 128 1023 0]]
	// 12345,67890 [[2 4 8 16 32] [64 128 256 512 1]]
	// 1(22.,33.)4(1,2) [[2 4 4 1023 16 2] [2 4 4 1023 16 4] [2 8 8 1023 16 2] [2 8 8 1023 16 4]]
	// 1(22.,33.)4(1) [[2 4 4 1023 16 2] [2 8 8 1023 16 2]]
	// 1(22.,33.)4() unexpected EOF
	// 1(22.,33.)4(1,) unexpected EOF
	// 1(22.,33.)4(,1) unexpected EOF
}

func TestParse(t *testing.T) {
	type args struct {
		r io.ByteReader
	}
	tests := []struct {
		name    string
		args    args
		want    [][]uint16
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			args:    args{bytes.NewBufferString("*")},
			want:    [][]uint16{{0}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readPattern() got = %v, want %v", got, tt.want)
			}
		})
	}
}
