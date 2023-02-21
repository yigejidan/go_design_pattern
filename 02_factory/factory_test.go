package _2_factory

import (
	"reflect"
	"testing"
)

func TestNewIRule(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRule
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: JsonRule{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: YamlRule{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRule(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKeyBoardFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IKeyBoardFactory
	}{
		{
			name: "HP",
			args: args{t: "HP"},
			want: HPKeyBoardFactory{},
		},
		{
			name: "Lenovo",
			args: args{t: "Lenovo"},
			want: LenovoKeyBoardFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeyBoardFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyBoardFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
