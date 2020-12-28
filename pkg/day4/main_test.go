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
			want: 2,
		},
		{
			name: "invalid passports",
			args: args{
				records: []string{
					"eyr:1972 cid:100",
					"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
					"",
					"iyr:2019",
					"hcl:#602927 eyr:1967 hgt:170cm",
					"ecl:grn pid:012533040 byr:1946",
					"",
					"hcl:dab227 iyr:2012",
					"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
					"",
					"hgt:59cm ecl:zzz",
					"eyr:2038 hcl:74454a iyr:2023",
					"pid:3556412378 byr:2007",
				},
			},
			want: 0,
		},
		{
			name: "valid passports",
			args: args{
				records: []string{
					"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
					"hcl:#623a2f",
					"",
					"eyr:2029 ecl:blu cid:129 byr:1989",
					"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
					"",
					"hcl:#888785",
					"hgt:164cm byr:2001 iyr:2015 cid:88",
					"pid:545766238 ecl:hzl",
					"eyr:2022",
					"",
					"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
				},
			},
			want: 4,
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
