package routree

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func Example_readPattern() {

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
		//
		"[1]*",
	} {
		pattern, err := readPattern(bytes.NewBufferString(s))
		if err != nil {
			fmt.Println(s, err)
			continue
		}
		fmt.Println(s, pattern)
	}

	// Output:
}

func Test_readPattern(t *testing.T) {
	type args struct {
		r io.ByteReader
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			args:    args{bytes.NewBufferString("*")},
			want:    []uint16{0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readPattern(tt.args.r)
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
