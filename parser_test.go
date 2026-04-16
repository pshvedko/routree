package routree

import (
	"io"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []Pattern
		err  error
	}{
		// TODO: Add test cases.
		{
			args: "*",
			err:  ErrIllegalSymbol{'*'},
		}, {
			args: ".*",
			want: []Pattern{{1023 | 0x8000}},
		}, {
			args: "1.*",
			want: []Pattern{{2, 1023 | 0x8000}},
		}, {
			args: "1**",
			err:  ErrIllegalSymbol{'*'},
		}, {
			args: "11*",
			want: []Pattern{{2, 2 | 0x8000}},
		}, {
			args: "11.",
			want: []Pattern{{2, 2, 1023}},
		}, {
			args: "1..",
			want: []Pattern{{2, 1023, 1023}},
		}, {
			args: "1.1",
			want: []Pattern{{2, 1023, 2}},
		}, {
			args: "[1]",
			want: []Pattern{{2}},
		}, {
			args: "[]",
			err:  ErrIllegalSymbol{']'},
		}, {
			args: "[0-1]",
			want: []Pattern{{3}},
		}, {
			args: "[0-2]",
			want: []Pattern{{7}},
		}, {
			args: "[0-9].",
			want: []Pattern{{1023, 1023}},
		}, {
			args: "[1-3]",
			want: []Pattern{{14}},
		}, {
			args: "[1-1]",
			err:  ErrIllegalRange{'1', '1'},
		}, {
			args: "[2-1]",
			err:  ErrIllegalRange{'2', '1'},
		}, {
			args: "[-1]",
			err:  ErrIllegalSymbol{'-'},
		}, {
			args: "[2-]",
			err:  ErrIllegalSymbol{']'},
		}, {
			args: "[1-3|7]",
			want: []Pattern{{142}},
		}, {
			args: "[1-3|7-9]",
			want: []Pattern{{910}},
		}, {
			args: "[1-3|7|9]",
			want: []Pattern{{654}},
		}, {
			args: "[1|2]",
			want: []Pattern{{6}},
		}, {
			args: "[1|2-4|0]",
			want: []Pattern{{31}},
		}, {
			args: "1[2|4]567.*",
			want: []Pattern{{2, 20, 32, 64, 128, 1023 | 0x8000}},
		}, {
			args: "12345,67890",
			want: []Pattern{{2, 4, 8, 16, 32}, {64, 128, 256, 512, 1}},
		}, {
			args: "1(22.,33.)4(1,2)",
			want: []Pattern{{2, 4, 4, 1023, 16, 2}, {2, 4, 4, 1023, 16, 4}, {2, 8, 8, 1023, 16, 2}, {2, 8, 8, 1023, 16, 4}},
		}, {
			args: "1(22.,33.)4(1)",
			want: []Pattern{{2, 4, 4, 1023, 16, 2}, {2, 8, 8, 1023, 16, 2}},
		}, {
			args: "1(22(1,2,4),3(3)3.)4",
			want: []Pattern{{2, 4, 4, 2, 16}, {2, 4, 4, 4, 16}, {2, 4, 4, 16, 16}, {2, 8, 8, 8, 1023, 16}},
		}, {
			args: "1(22.,33.)4()",
			err:  io.ErrUnexpectedEOF,
		}, {
			args: "1(22.,33.)4(1,)",
			err:  io.ErrUnexpectedEOF,
		}, {
			args: "1(22.,33.)4(,1)",
			err:  io.ErrUnexpectedEOF,
		}, {
			args: "123#.*",
			want: []Pattern{{2, 4, 8 | 0x4000, 1023 | 0x8000}},
		}, {
			args: "123##.*",
			err:  ErrIllegalSymbol{'#'},
		}, {
			args: "123#*",
			want: []Pattern{{2, 4, 8 | 0x4000 | 0x8000}},
		}, {
			args: "[1|23]",
			err:  ErrIllegalSymbol{'3'},
		}, {
			args: "[1-23]",
			err:  ErrIllegalSymbol{'3'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePattern(tt.args)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("ParsePattern() error = %v, err %v", err, tt.err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePattern() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePhone(t *testing.T) {
	type args struct {
		number  string
		options []Option
	}
	tests := []struct {
		name string
		args args
		want Pattern
		err  error
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				number: "79991234567",
			},
			want: Pattern{128, 512, 512, 512, 2, 4, 8, 16, 32, 64, 128},
		},
		{
			name: "",
			args: args{
				number:  "+79991234567",
				options: []Option{E164},
			},
			want: Pattern{128, 512, 512, 512, 2, 4, 8, 16, 32, 64, 128},
		},
		{
			name: "",
			args: args{
				number:  "+",
				options: []Option{E164},
			},
			err: io.ErrUnexpectedEOF,
		},
		{
			name: "",
			args: args{
				options: []Option{ErrFormat(0)},
			},
			err: ErrFormat(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePhone(tt.args.number, tt.args.options...)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("ParsePhone() error = %v, err %v", err, tt.err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePhone() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
