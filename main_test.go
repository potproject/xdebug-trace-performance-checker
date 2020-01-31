package main

import (
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
