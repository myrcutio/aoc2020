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
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "supplied test case",
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
			},
			want: "...",
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
