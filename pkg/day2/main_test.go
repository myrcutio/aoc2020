package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_iteratePasswords(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "supplied test case",
			args: args{
				records: [][]string{{
					"1-3 a: abcde",
				},
					{
						"1-3 b: cdefg",
					},
					{
						"2-9 c: ccccccccc",
					}},
			},
			want: 2,
		},
		{
			name: "testData",
			args: args{
				records: func() [][]string {
					csvFile, _ := filepath.Abs("./input.csv")
					in, err := os.Open(csvFile)
					if err != nil {
						panic(err)
					}

					r := csv.NewReader(in)

					records, err := r.ReadAll()
					if err != nil {
						log.Fatal(err)
					}
					return records
				}(),
			},
			want: 434,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iteratePasswords(tt.args.records); got != tt.want {
				t.Errorf("iteratePasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newIteratePasswords(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "supplied test case",
			args: args{
				records: [][]string{{
						"1-3 a: abcde",
					},
					{
						"1-3 b: cdefg",
					},
					{
						"2-9 c: ccccccccc",
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newIteratePasswords(tt.args.records); got != tt.want {
				t.Errorf("newIteratePasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}