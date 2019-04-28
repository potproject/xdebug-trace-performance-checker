package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var fp *os.File
	var err error
	if len(os.Args) < 3 {
		panic("invaild args.")
	}

	// filename
	fileName := os.Args[1]

	// times
	var times float64
	times, err = strconv.ParseFloat(os.Args[2], 64)

	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	lines := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		lines = append(lines, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var beforeTimes float64
	for i, word := range lines {
		if i == 0 {
			continue
		}
		// group[1] = 時間
		// group[2] = メモリ使用量
		// group[3] = 空白
		// group[4] = メソッド
		assined := regexp.MustCompile(`\s+(\d+\.\d+)\s+(\d+)(\s+)(.+)`)
		group := assined.FindSubmatch([]byte(word))
		if len(group) < 5 {
			continue
		}
		afterTimes, err := strconv.ParseFloat(string(group[1]), 64)
		if err != nil {
			panic(err)
		}
		if (beforeTimes + times) < afterTimes {
			fmt.Println("Line:", i, " Duration:", (afterTimes - beforeTimes))
			fmt.Println(getStacktrace(lines, i, len(group[3])))
		}
		beforeTimes = afterTimes
	}
}

func getStacktrace(lines []string, i int, space int) string {
	result := fmt.Sprintln(lines[i])
	for i > 0 {
		i--
		assined := regexp.MustCompile(`\s+(\d+\.\d+)\s+(\d+)(\s+)(.+)`)
		group := assined.FindSubmatch([]byte(lines[i]))
		if len(group) < 5 {
			continue
		}
		if len(group[3]) < space {
			result = result + fmt.Sprintln(lines[i])
			space = len(group[3])
		}
	}
	return result
}
