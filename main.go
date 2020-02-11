package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Trace Data
type Trace struct {
	time       float64 // 全体処理時間
	usedMemory int64   // 全体使用しているメモリ
	depth      int     // 変数の深さ。rootの場合は1
	method     string  // メソッド情報
	filePath   string  // ファイルパス
	line       int64   // 行数
}

func main() {
	fileName, times, err := getParams(os.Args)
	if err != nil {
		panic(err)
	}

	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	lines, err := getLines(fp)
	if err != nil {
		panic(err)
	}

	traces := parseTrace(lines)

	var beforeTimes float64
	for i, trace := range traces {
		if (beforeTimes + times) < trace.time {
			fmt.Println("Line:", i+2, " Duration:", (trace.time - beforeTimes))
			sts := getStacktrace(traces, i)
			for _, st := range sts {
				fmt.Printf("%v\n", st)
			}
			fmt.Println()
		}
		beforeTimes = trace.time
	}
}

func parseTrace(lines []string) (traces []Trace) {
	for i, word := range lines {
		if i == 0 {
			continue
		}
		// group[1] = 時間
		// group[2] = メモリ使用量
		// group[3] = 空白 初回:3 で以降2ずつ増える
		// group[4] = メソッド
		// group[5] = ファイルパス
		// group[6] = 行数
		assined := regexp.MustCompile(`\s+(\d+\.\d+)\s+(\d+)(\s+)->\s(\S+)\s([^:]+):(\d+)`)
		group := assined.FindSubmatch([]byte(word))
		if len(group) < 5 {
			continue
		}
		time, _ := strconv.ParseFloat(string(group[1]), 64)
		usedMemory, _ := strconv.ParseInt(string(group[2]), 10, 64)
		depth := depthCount(string(group[3]))
		method := string(group[4])
		filePath := string(group[5])
		line, _ := strconv.ParseInt(string(group[6]), 10, 64)
		t := Trace{
			time:       time,
			usedMemory: usedMemory,
			depth:      depth,
			method:     method,
			filePath:   filePath,
			line:       line,
		}
		traces = append(traces, t)
	}
	return
}

func depthCount(space string) int {
	// 3つのスペースはデフォルトで入るため除外
	// 1つの深さにより2つのスペースを使用する
	return (strings.Count(space, " ") - 3) / 2
}

func depthPadding(depth int) string {
	return strings.Repeat(" ", 3+(depth*2))
}

// getParams
// Args[1]: fileName
// Args[2]: times
func getParams(args []string) (fileName string, times float64, err error) {
	// filename
	fileName = args[1]

	// times
	times, err = strconv.ParseFloat(args[2], 64)
	return
}

func getLines(fp *os.File) (lines []string, err error) {
	scanner := bufio.NewScanner(fp)
	lines = []string{}
	for scanner.Scan() {
		row := scanner.Text()
		lines = append(lines, row)
	}
	return

}
func getStacktrace(traces []Trace, index int) (stackTraces []Trace) {
	stackTraces = append(stackTraces, traces[index])
	beforeDepth := traces[index].depth
	for i := index; i > -1 && beforeDepth > 0; i-- {
		if traces[i].depth < beforeDepth {
			stackTraces = append(stackTraces, traces[i])
			beforeDepth = traces[i].depth
		}
	}
	return
}
