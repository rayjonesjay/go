package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "foo" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be foo", 9, "foo"},
		{"3 should be foo", 3, "foo"},
		{"4 should not foo", 4, "4"},
		{"0 should be foo", 0, "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got %q want %q", ans, tt.want)
			}
		})
	}
}

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(input)
	//t.Log(result)
	//t.Logf("->>>%s<<<<-", result)
	if result != "f oo" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "foo")
	}
	t.Fatal("stop the test now, we have seen enough")
	t.Errorf("this wont be executed")
}

type tests struct {
	name  string
	input int
	want  string
}

func TestFooerParallel(t *testing.T) {
	var ts = []tests{
		{"test 3 in parallel", 3, "foo"},
		{"test 7 in parallel", 7, "7"},
	}

	for _, test := range ts {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := Fooer(test.input)
			if test.want != result {
				t.Errorf("got %q want %q", result, test.want)
			}
		})
	}

	//t.Run("test 3 in parallel", func(t *testing.T) {
	//	t.Parallel()
	//	result := Fooer(3)
	//	if result != "foo" {
	//		t.Errorf("result was incorrect, got: %s, want: %s.", result, "foo")
	//	}
	//})
	//t.Run("test 7 in parallel", func(t *testing.T) {
	//	t.Parallel()
	//	result := Fooer(7)
	//	if result != "7" {
	//		t.Errorf("result was incorrect, got: %s, want: %s.", result, "7")
	//	}
	//})
}
