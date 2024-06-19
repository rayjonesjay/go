package caret

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	type args struct {
		a Caret
		b Caret
	}

	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}
	c := []string{"ad", "be", "cf"}
	tests := []struct {
		name string
		args args
		want Caret
	}{
		{
			name: "",
			args: args{
				a: a,
				b: b,
			},
			want: c,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Append(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Append() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		c Caret
	}
	tests := []struct {
		name string
		args args
		want Caret
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Copy(tt.args.c); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Copy() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewCaret(t *testing.T) {
	tests := []struct {
		name string
		want Caret
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewCaret(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewCaret() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNilOrEmpty(t *testing.T) {
	type args struct {
		caret []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NilOrEmpty(tt.args.caret); got != tt.want {
					t.Errorf("NilOrEmpty() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
