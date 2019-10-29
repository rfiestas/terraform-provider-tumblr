package tumblr

import (
	"fmt"
	"reflect"
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
			name: "{}",
			args: args{
				v: "queue",
			},
			wantWs: nil,
			wantEs: nil,
		},
		{
			name: "{}",
			args: args{
				v: "novalid",
			},
			wantEs: []error{fmt.Errorf("State 'novalid' is not valid. Choose one of these: [private draft queue published]")},
		},
		{
			name: "{}",
			args: args{
				v: 1,
			},
			wantEs: []error{fmt.Errorf("Expected name to be string")},
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
			name: "{}",
			args: args{
				v: "2019-10-26 21:39:29 GTM",
			},
			wantWs: nil,
			wantEs: nil,
		},
		{
			name: "{}",
			args: args{
				v: "2019-Jan-26 21:39:29 GTM",
			},
			wantEs: []error{fmt.Errorf("Date '2019-Jan-26 21:39:29 GTM' is not valid format. Format must be '2006-01-02 15:04:05 MST'")},
		},
		{
			name: "{}",
			args: args{
				v: 1,
			},
			wantEs: []error{fmt.Errorf("Expected name to be string")},
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
