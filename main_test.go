package main

import (
	"testing"

	"golang.org/x/tools/benchmark/parse"
)

var bTests = []struct {
	line    string // input
	name    string // expected result
	arg     string
	nsperop float64
}{
	{"BenchmarkF2_F0000000-4		50000000	        29.4 ns/op", "F2", "F0000000", 29.4},
	{"BenchmarkF0_FF-2		10000000	        37.4 ns/op", "F0", "FF", 37.4},
	{"BenchmarkF_0-2		40000000	        11.2 ns/op", "F", "0", 11.2},
	{"BenchmarkF3/quicksort_100-4		40000000	        11.2 ns/op", "F3/quicksort", "100", 11.2},
}

func TestParser(t *testing.T) {
	for _, tt := range bTests {
		b, _ := parse.ParseLine(tt.line)
		name, arg, _, _ := parseNameArgThread(b.Name)
		if name != tt.name {
			t.Errorf("parseNameArgThread(%s): expected %s, actual %s", b.Name, tt.name, name)
		}
		if arg != tt.arg {
			t.Errorf("parseNameArgThread(%s): expected %s, actual %s", b.Name, tt.arg, arg)
		}
		if b.NsPerOp != tt.nsperop {
			t.Errorf("parseNameArgThread(%s): expected %f, actual %f", b.Name, tt.nsperop, b.NsPerOp)
		}
	}
}
