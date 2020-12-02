package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Password struct {
	min      int
	max      int
	letter   string
	sequence string
}

//Day2 is...
func Day2() {
	data, _ := readFilePass("data_day2.txt")
	passwords := makeStructs(data)
	iteratePasswords1(passwords)
	iteratePasswords2(passwords)

}

func readFilePass(fname string) (str []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n")
	str = make([]string, 0, len(lines))
	//fmt.Println(lines)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		str = append(str, l)
	}

	return str, nil
}

func makeStructs(data []string) (passwords []*Password) {
	passwords = []*Password{}
	for _, el := range data {
		str := strings.Split(el, " ")
		nums := strings.Split(str[0], "-")
		min, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err, nums[0])
		}
		max, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err, nums[1])
		}
		letter := strings.Split(str[1], ":")[0]
		sequence := (str[2])
		password := &Password{min, max, letter, sequence}
		passwords = append(passwords, password)
		//fmt.Println(password)
	}

	return passwords
}

func iteratePasswords1(data []*Password) {
	passQty := 0
	for _, el := range data {
		qty := strings.Count(el.sequence, el.letter)
		if qty >= el.min && qty <= el.max {
			passQty++
		}
	}
	fmt.Println(passQty)
}

func iteratePasswords2(data []*Password) {
	passQty := 0
	for _, el := range data {
		first := string(el.sequence[el.min-1]) == el.letter
		fmt.Println(first)

		second := string(el.sequence[el.max-1]) == el.letter
		if first != second {
			passQty++
		}
	}
	fmt.Println(passQty)
}
