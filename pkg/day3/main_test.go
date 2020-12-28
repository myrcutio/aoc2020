package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_operation(t *testing.T) {
	type args struct {
		records [][]string
		slope  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "slope of 3,1",
			args: args{
				records: func() [][]string {
					csvFile, _ := filepath.Abs("./test.csv")
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
				slope: []int{3,1},
			},
			want: 7,
		},
		{
			name: "slope of 1,1",
			args: args{
				records: func() [][]string {
					csvFile, _ := filepath.Abs("./test.csv")
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
				slope: []int{1,1},
			},
			want: 2,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := operation(tt.args.records, tt.args.slope); got != tt.want {
				t.Errorf("operation() = %v, want %v", got, tt.want)
			}
		})
	}
}
