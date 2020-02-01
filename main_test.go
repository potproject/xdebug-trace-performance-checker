package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetParams(t *testing.T) {
	var expect []string
	expect = []string{"", "trace.xt", "0.1"}
	actualFileName := "trace.xt"
	actualTimes := 0.1
	fileName, times, err := getParams(expect)
	if err != nil {
		t.Error(err)
	}
	if fileName != actualFileName {
		t.Errorf("invalid fileName")
	}
	if times != actualTimes {
		t.Errorf("invalid times")
	}
}

func TestDepth(t *testing.T) {
	s := " "
	if depthCount(strings.Repeat(s, 5)) != 1 {
		t.Errorf("invalid depthCount")
	}
	if depthCount(strings.Repeat(s, 23)) != 10 {
		t.Errorf("invalid depthCount")
	}
	if depthPadding(1) != strings.Repeat(s, 5) {
		t.Errorf("invalid depthPadding")
	}
	if depthPadding(10) != strings.Repeat(s, 23) {
		t.Errorf("invalid depthPadding")
	}
}

func TestParseTraceSuccess(t *testing.T) {
	expect := []string{
		"TRACE START [2020-02-01 12:30:32]",
		"    0.0050      68952   -> {main}() /var/www/test.php:0",
		"    0.0100     240056     -> include(/var/www/test2.php) /var/www/test.php:3",
		"    0.0295     409624     -> test() /var/www/test.php:6",
		"    0.1962     410440       -> foo() /var/www/test2.php:10",
		"    0.3201     410472     -> test2() /var/www/test.php:7",
		"    0.3211     410568       -> bar() /var/www/test2.php:24",
	}
	actual := []Trace{
		Trace{time: 0.005, usedMemory: 68952, depth: 0, method: "{main}()", filePath: "/var/www/test.php", line: 0},
		Trace{time: 0.01, usedMemory: 240056, depth: 1, method: "include(/var/www/test2.php)", filePath: "/var/www/test.php", line: 3},
		Trace{time: 0.0295, usedMemory: 409624, depth: 1, method: "test()", filePath: "/var/www/test.php", line: 6},
		Trace{time: 0.1962, usedMemory: 410440, depth: 2, method: "foo()", filePath: "/var/www/test2.php", line: 10},
		Trace{time: 0.3201, usedMemory: 410472, depth: 1, method: "test2()", filePath: "/var/www/test.php", line: 7},
		Trace{time: 0.3211, usedMemory: 410568, depth: 2, method: "bar()", filePath: "/var/www/test2.php", line: 24}}
	result := parseTrace(expect)
	if !reflect.DeepEqual(result, actual) {
		t.Errorf("parseTrace() = %v, want %v", result, actual)
	}
}

func TestParseTraceIrregular(t *testing.T) {
	expect := []string{
		"TRACE START [2020-02-01 12:30:32]",
		"    0.0050      68952   -> {main}() /var/www/test.php:0",
		" invalid string",
		"    0.0100     240056     -> include(/var/www/test2.php) /var/www/test.php:3",
		"    0.0295     409624     -> test() /var/www/test.php:6",
		"    0.1962     410440       -> foo() /var/www/test2.php:10",
		"    0.3201     410472     -> test2() /var/www/test.php:7",
		"    0.3211     410568       -> bar() /var/www/test2.php:24",
		"",
	}
	actual := []Trace{
		Trace{time: 0.005, usedMemory: 68952, depth: 0, method: "{main}()", filePath: "/var/www/test.php", line: 0},
		Trace{time: 0.01, usedMemory: 240056, depth: 1, method: "include(/var/www/test2.php)", filePath: "/var/www/test.php", line: 3},
		Trace{time: 0.0295, usedMemory: 409624, depth: 1, method: "test()", filePath: "/var/www/test.php", line: 6},
		Trace{time: 0.1962, usedMemory: 410440, depth: 2, method: "foo()", filePath: "/var/www/test2.php", line: 10},
		Trace{time: 0.3201, usedMemory: 410472, depth: 1, method: "test2()", filePath: "/var/www/test.php", line: 7},
		Trace{time: 0.3211, usedMemory: 410568, depth: 2, method: "bar()", filePath: "/var/www/test2.php", line: 24}}
	result := parseTrace(expect)
	if !reflect.DeepEqual(result, actual) {
		t.Errorf("parseTrace() = %v, want %v", result, actual)
	}
}
