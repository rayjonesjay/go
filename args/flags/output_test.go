package flags

import (
	"testing"
)

func TestInspectFlagAndFile(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Correct arg", args{[]string{"--output=file.txt"}}, "file.txt"},
		{"Correct arg alphanumerical", args{[]string{"--output=544file1234.txt"}}, "544file1234.txt"},
		{"Correct arg (dots file)", args{[]string{"--output=..."}}, "..."},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := InspectFlagAndFile(tt.args.args); got != tt.want {
					t.Errorf("InspectFlagAndFile() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
