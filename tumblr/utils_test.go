package tumblr

import (
	"testing"
)

func Test_stringToUint(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "string(int) to int",
			args: args{
				str: "1234567890",
			},
			want: 1234567890,
		}, {
			name: "string(string) to int",
			args: args{
				str: "numbers",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToUint(tt.args.str); got != tt.want {
				t.Errorf("stringToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uintToString(t *testing.T) {
	type args struct {
		integer uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "int into string",
			args: args{
				integer: 1234567890,
			},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uintToString(tt.args.integer); got != tt.want {
				t.Errorf("uintToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToMd5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string to md5",
			args: args{
				str: "1234567890",
			},
			want: "e807f1fcf82d132f9bb018ca6738a19f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToMd5(tt.args.str); got != tt.want {
				t.Errorf("stringToMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "To camel case",
			args: args{
				str: "foo_var_foo_var",
			},
			want: "FooVarFooVar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCamelCase(tt.args.str); got != tt.want {
				t.Errorf("toCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
