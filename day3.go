package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Day3 is...
func Day3() {
	data, _ := readFilePath("data_day3.txt")
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	treesSlice := []int{}
	for i := range slopes {
		trees := moving(data, slopes[i][0], slopes[i][1])
		treesSlice = append(treesSlice, trees)
	}
	result := 1

	for _, tree := range treesSlice {
		result *= tree
	}
	fmt.Println(treesSlice)
	fmt.Println(result)

}

func readFilePath(fname string) (maps [][]byte, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n")
	maps = make([][]byte, 0)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		str := []byte(l)
		maps = append(maps, str)
	}

	return maps, nil
}

func moving(data [][]byte, right int, down int) int {
	trees := 0
	row := 0
	column := 0
	for {
		if column >= len(data[0]) {
			column = column % len(data[0])
		}
		if row >= len(data) {
			break
		}
		if data[row][column] == 35 {
			trees++
		}
		column = column + right
		row = row + down

	}
	return trees
}
