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
		preambleSize int
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
				preambleSize: 5,
			},
			want: 62,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operation(tt.args.records, tt.args.preambleSize); got != tt.want {
				t.Errorf("operation() = %v, want %v", got, tt.want)
			}
		})
	}
}
