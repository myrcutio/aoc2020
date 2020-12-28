package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	BirthYear string `csv:"byr"`
	IssueYear string `csv:"iyr"`
	ExpirationYear string `csv:"eyr"`
	Height string `csv:"hgt"`
	HairColor string `csv:"hcl"`
	EyeColor string `csv:"ecl"`
	PassportID string `csv:"pid"`
	CountryID string `csv:"cid"`
}

func operation(records []string) (int, error) {

	var passports []passport

	rowsRemain := true
	for rowsRemain {
		raw, remaining := identifyPassportRows(records)
		if remaining == nil {
			rowsRemain = false
		} else {
			records = remaining
		}
		if raw != "" {
			passports = append(passports, parsePassport(raw))
		}
	}

	validCount := 0
	for _, v := range passports {
		if isValidPassport(v) {
			validCount++
		}
	}

	return validCount, nil
}

func identifyPassportRows(rows []string) (string, []string) {
	raw := ""
	if rows[0] == "" {
		return raw, rows[1:]
	}

	foundRows := []string{}
	for i, value := range rows {
		v := value

		if v == "" {
			for _, r := range foundRows {
				raw = raw + " " + r
			}
			remainingRows := rows[i:]
			return strings.Trim(raw, " "), remainingRows
		} else {
			foundRows = append(foundRows, v)
		}
	}
	for _, r := range foundRows {
		raw = raw + " " + r
	}

	return strings.Trim(raw, " "), nil
}

func parsePassport(raw string) passport {
	p := passport{}
	rawFields := strings.Split(raw, " ")

	for _, v := range rawFields {
		field := strings.Split(v, ":")
		switch field[0] {
		case "byr":
			p.BirthYear = field[1]
		case "iyr":
			p.IssueYear = field[1]
		case "eyr":
			p.ExpirationYear = field[1]
		case "hgt":
			p.Height = field[1]
		case "hcl":
			p.HairColor = field[1]
		case "ecl":
			p.EyeColor = field[1]
		case "pid":
			p.PassportID = field[1]
		case "cid":
			p.CountryID = field[1]
		}
	}

	return p
}

func isValidPassport(p passport) bool {
	validBYear := validIntRange(p.BirthYear, 1920, 2002)
	if p.BirthYear == "" || !validBYear {
		return false
	}
	validIYear := validIntRange(p.IssueYear, 2010, 2020)
	if p.IssueYear == "" || !validIYear {
		return false
	}
	validEYear := validIntRange(p.ExpirationYear, 2020, 2030)
	if p.ExpirationYear == "" || !validEYear {
		return false
	}
	validHeight := validHeight(p.Height)
	if p.Height == "" || !validHeight {
		return false
	}
	validHair := validHairColor(p.HairColor)
	if p.HairColor == "" || !validHair {
		return false
	}
	validEye := validEyeColor(p.EyeColor)
	if p.EyeColor == "" || !validEye {
		return false
	}
	validId := validID(p.PassportID)
	if p.PassportID == "" || !validId {
		return false
	}

	return true
}

func validIntRange(v string, min, max int64) bool {
	by, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return false
	}
	return by >= min && by <= max
}

func validHeight(height string) bool {
	r := regexp.MustCompile(`^(?P<measure>\d{2,3})(?P<unit>\w{2})$`)
	values := r.FindStringSubmatch(height)
	if len(values) < 2 {
		return false
	}
	if values[2] == "cm" {
		h, err := strconv.Atoi(values[1])
		if err != nil {
			return false
		}
		return h >= 150 && h <= 193
	}
	if values[2] == "in" {
		h, err := strconv.Atoi(values[1])
		if err != nil {
			return false
		}
		return h >= 59 && h <= 76
	}
	return false
}

func validHairColor(color string) bool {
	r := regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	return r.Match([]byte(color))
}

func validEyeColor(color string) bool {
	validColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return validColors[color]
}

func validID(id string) bool {
	r := regexp.MustCompile(`^[0-9]{9}$`)
	return r.Match([]byte(id))
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day4/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer, _ := operation(records)

	fmt.Printf("current data transform: %d", answer)
}