package tumblr

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_validateState(t *testing.T) {
	type args struct {
		v interface{}
		k string
	}
	tests := []struct {
		name   string
		args   args
		wantWs []string
		wantEs []error
	}{
		{
			name: "Valid State",
			args: args{
				v: "queue",
			},
			wantWs: nil,
			wantEs: nil,
		},
		{
			name: "Invalid State",
			args: args{
				v: "novalid",
			},
			wantEs: []error{fmt.Errorf("State 'novalid' is not valid. Choose one of these: [private draft queue published]")},
		},
		{
			name: "State is not a string",
			args: args{
				v: 1,
			},
			wantEs: []error{fmt.Errorf(errorString)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWs, gotEs := validateState(tt.args.v, tt.args.k)
			if !reflect.DeepEqual(gotWs, tt.wantWs) {
				t.Errorf("validateState() gotWs = %v, want %v", gotWs, tt.wantWs)
			}
			if !reflect.DeepEqual(gotEs, tt.wantEs) {
				t.Errorf("validateState() gotEs = %v, want %v", gotEs, tt.wantEs)
			}
		})
	}
}

func Test_validateDate(t *testing.T) {
	type args struct {
		v interface{}
		k string
	}
	tests := []struct {
		name   string
		args   args
		wantWs []string
		wantEs []error
	}{
		{
			name: "Valid Date",
			args: args{
				v: "2019-10-26 21:39:29 GTM",
			},
			wantWs: nil,
			wantEs: nil,
		},
		{
			name: "Invalid Date",
			args: args{
				v: "2019-Jan-26 21:39:29 GTM",
			},
			wantEs: []error{fmt.Errorf("Date '2019-Jan-26 21:39:29 GTM' is not valid format. Format must be '2006-01-02 15:04:05 MST'")},
		},
		{
			name: "Date is not string",
			args: args{
				v: 1,
			},
			wantEs: []error{fmt.Errorf(errorString)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWs, gotEs := validateDate(tt.args.v, tt.args.k)
			if !reflect.DeepEqual(gotWs, tt.wantWs) {
				t.Errorf("validateState() gotWs = %v, want %v", gotWs, tt.wantWs)
			}
			if !reflect.DeepEqual(gotEs, tt.wantEs) {
				t.Errorf("validateState() gotEs = %v, want %v", gotEs, tt.wantEs)
			}
		})
	}
}

func Test_validateData64(t *testing.T) {
	type args struct {
		v interface{}
		k string
	}
	tests := []struct {
		name   string
		args   args
		wantWs []string
		wantEs []error
	}{
		{
			name: "File in Data64 exist",
			args: args{
				v: "existing_file.png", // ¯\_(ツ)_/¯
			},
			wantWs: nil,
			wantEs: nil,
		},
		{
			name: "File in Data64, doesn't exist",
			args: args{
				v: "no_file.png",
			},
			wantWs: nil,
			wantEs: []error{fmt.Errorf("File 'no_file.png' doesn't exist")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.v.(string) == "existing_file.png" {
				_ = ioutil.WriteFile(tt.args.v.(string), []byte("Hello"), 0755)
				defer os.Remove(tt.args.v.(string))
			}
			gotWs, gotEs := validateData64(tt.args.v, tt.args.k)
			if !reflect.DeepEqual(gotWs, tt.wantWs) {
				t.Errorf("validateState() gotWs = %v, want %v", gotWs, tt.wantWs)
			}
			if !reflect.DeepEqual(gotEs, tt.wantEs) {
				t.Errorf("validateState() gotEs = %v, want %v", gotEs, tt.wantEs)
			}
		})
	}
	testsContains := []struct {
		name   string
		args   args
		wantWs []string
		wantEs string
	}{
		{
			name: "Filename in Data64 too long",
			args: args{
				v: strings.Repeat("X", 257),
			},
			wantWs: nil,
			wantEs: "file name too long",
		},
		{
			name: "Data64 is not string",
			args: args{
				v: 1,
			},
			wantWs: nil,
			wantEs: errorString,
		},
	}
	for _, tt := range testsContains {
		t.Run(tt.name, func(t *testing.T) {
			gotWs, gotEs := validateData64(tt.args.v, tt.args.k)

			if !reflect.DeepEqual(gotWs, tt.wantWs) {
				t.Errorf("validateState() gotWs = %v, want %v", gotWs, tt.wantWs)
			}
			if !strings.Contains(gotEs[0].Error(), tt.wantEs) {
				t.Errorf("validateState() gotEs = %v, want %v", gotEs, tt.wantEs)
			}
		})
	}
}
