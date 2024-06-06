package fmtx

import "testing"

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}

	i := 1000
	str := "one thousand"
	tests := []struct {
		name string
		args args
	}{
		{"Single line error", args{format: "test message\n"}},
		{
			name: "Multi-line error", args: args{
				format: "test message\nmessage #%d\n",
				args:   []interface{}{i},
			},
		},

		{
			name: "Single line error (int)", args: args{
				format: "test message #%d\n",
				args:   []interface{}{i},
			},
		},

		{
			name: "Single line error (int, string)", args: args{
				format: "test message #%d -> %s\n",
				args:   []interface{}{i, str},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				Errorf(tt.args.format, tt.args.args...)
			},
		)
	}
}
