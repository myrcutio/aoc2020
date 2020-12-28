package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func Test_operation(t *testing.T) {
	type args struct {
		records []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "supplied test case",
			args: args{
				records: func() []string {
					inputFile, _ := filepath.Abs("test.csv")
					r, _ := ioutil.ReadFile(inputFile)

					records := strings.Split(string(r), "\n")
					return records
				}(),
			},
			want: 6,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := operation(tt.args.records); got != tt.want {
				t.Errorf("operation() = %v, want %v", got, tt.want)
			}
		})
	}
}
