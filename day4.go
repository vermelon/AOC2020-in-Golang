package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//Day4 is...
func Day4() {
	data, _ := readFileBatch("data_day4.txt")
	validData := checkValidity(data)
	checkValidity2(validData)
}

func readFileBatch(fname string) (str []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n\r\n")
	str = make([]string, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		str = append(str, l)
	}

	return str, nil
}

func checkValidity(data []string) (validData []string) {
	validData = []string{}
	for _, el := range data {
		if strings.Contains(el, "byr") && strings.Contains(el, "iyr") && strings.Contains(el, "eyr") && strings.Contains(el, "hgt") && strings.Contains(el, "hcl") && strings.Contains(el, "ecl") && strings.Contains(el, "pid") {
			validData = append(validData, el)
		}
	}
	return validData
}

func checkValidity2(data []string) {
	validPasssports := 0
	string1 := []string{}
	for _, el := range data {
		str := strings.ReplaceAll(el, "\r\n", " ")
		string1 = append(string1, str)
	}
	for _, eachLine := range string1 {
		validFields := 0
		newMap := make(map[string]string)
		keyValuePairs := strings.Split(eachLine, " ")
		for _, keyValuePair := range keyValuePairs {
			keyValuePair := strings.Split(keyValuePair, ":")
			newMap[keyValuePair[0]] = keyValuePair[1]
		}
		byr, _ := strconv.Atoi(newMap["byr"])
		iyr, _ := strconv.Atoi(newMap["iyr"])
		eyr, _ := strconv.Atoi(newMap["eyr"])
		hgt := newMap["hgt"]
		hcl := newMap["hcl"]
		ecl := newMap["ecl"]
		pid := newMap["pid"]

		if byr >= 1920 && byr <= 2002 {
			validFields++
		}
		if iyr >= 2010 && iyr <= 2020 {
			validFields++
		}
		if eyr >= 2020 && eyr <= 2030 {
			validFields++
		}
		if strings.Contains(hgt, "in") {
			hgtIn, _ := strconv.Atoi(hgt[:len(hgt)-2])
			if hgtIn >= 59 && hgtIn <= 76 {
				validFields++
			}
		}
		if strings.Contains(hgt, "cm") {
			hgtCm, _ := strconv.Atoi(hgt[:len(hgt)-2])
			if hgtCm >= 150 && hgtCm <= 193 {
				validFields++
			}
		}
		validHairColor := regexp.MustCompile(`^#([a-fA-F0-9]){6}`)
		if validHairColor.MatchString(hcl) {
			validFields++
		}
		if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
			validFields++
		}
		if len(pid) == 9 {
			validFields++
		}
		if validFields == 7 {
			validPasssports++

		}

	}
	fmt.Println(validPasssports)
}
