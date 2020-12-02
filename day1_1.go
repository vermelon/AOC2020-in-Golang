package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

//Day1 is...
func Day1() {
	data, _ := readFile("data_day1_1.txt")
	findSumTwo(data)
	findSumThree(data)

}

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func findSumTwo(data []int) {
	sort.Ints(data)
	for _, v := range data {
		v2 := 2020 - v
		if j := sort.SearchInts(data, v2); v2 == data[j] {
			fmt.Println(v, v2)
			fmt.Println(v * v2)
			return
		}
	}
}

func findSumThree(data []int) {
	sort.Ints(data)
	for i, v := range data {
		//v2 := 2020 - v
		ind1 := i + 1
		ind2 := len(data) - 1
		for ind1 < ind2 {
			if v+data[ind1]+data[ind2] < 2020 {
				//fmt.Println(v + data[ind1] + data[ind2])
				ind1++
				//fmt.Println("ind1 is")
				//fmt.Println(ind1)
			}
			if v+data[ind1]+data[ind2] > 2020 {
				//fmt.Println(v + data[ind1] + data[ind2])
				ind2 = ind2 - 1
				//fmt.Println("ind2 is")
				//fmt.Println(ind2)
			}
			if v+data[ind1]+data[ind2] == 2020 {
				//fmt.Println(v + data[ind1] + data[ind2])
				fmt.Println(ind1, ind2, v)
				fmt.Println(data[ind1], data[ind2], v)
				fmt.Println(data[ind1] * data[ind2] * v)
				return
			}
		}
	}
}
