package flags

import (
	"testing"
)

func TestMoreFlags(t *testing.T) {

	mArgs := []string{
		"--color=red",                // 0
		"hello",                      // 1
		"--output=data.txt",          // 2
		"--color=blue",               // 3
		"-color=blue",                // 4
		"world",                      // 5
		"hello world! this is hard!", // 6
		"shadow",                     // 7
	}

	mArgs1 := []string{
		"--color=red",                // 0
		"hello",                      // 1
		"--output=data.txt",          // 2
		"--color=blue",               // 3
		"--",                         // 4
		"world",                      // 5
		"hello world! this is hard!", // 6
		"shadow",                     // 7
	}

	mArgs2 := []string{
		"--color=red",                // 0
		"hello",                      // 1
		"--",                         // 2
		"--output=data.txt",          // 3
		"--color=blue",               // 4
		"--",                         // 5
		"world",                      // 6
		"hello world! this is hard!", // 7
		"shadow",                     // 8
	}

	type args struct {
		args []string
		i    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Non-existence index", args{mArgs, -1}, false},
		{"No args", args{nil, 0}, false},
		{"No args + Non-existence index", args{nil, -10}, false},
		{"", args{mArgs, 0}, true},
		{"", args{mArgs, 1}, true},
		{"", args{mArgs, 2}, true},
		{"", args{mArgs, 3}, true},
		{"", args{mArgs, 4}, false},
		{"", args{mArgs, 5}, false},
		{"", args{mArgs, 6}, false},
		{"Non-existence index (over size)", args{mArgs, 7}, false},

		{"Nil flag (--)", args{mArgs1, 4}, true},
		{"Nil flag control", args{mArgs1, 5}, false},

		{"End flag (--)", args{mArgs2, 0}, true},
		{"End flag (--)", args{mArgs2, 1}, true},
		{"End flag (--)", args{mArgs2, 2}, true},
		{"End flag (--)", args{mArgs2, 3}, false},
		{"End flag (--)", args{mArgs2, 4}, false},
		{"End flag (--)", args{mArgs2, 5}, false},
		{"End flag (--)", args{mArgs2, 6}, false},
		{"End flag (--)", args{mArgs2, 7}, false},
		{"End flag (--)", args{mArgs2, 8}, false},
		{"End flag (--)", args{mArgs2, 100}, false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := MoreFlags(tt.args.args, tt.args.i); got != tt.want {
					t.Errorf("MoreFlags() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
