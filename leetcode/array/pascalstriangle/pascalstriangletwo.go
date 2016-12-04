package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

// reverse reverses a slice of integers.
func reverse(r []int) []int {
	n, h := len(r), len(r)/2
	for i := 0; i < h; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return r
}

func main() {
	flag.Parse()
	//Setup  profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

	m := 5 //max iterations
	prev := []int{1}
	fmt.Println(prev)

	for row := 1; row < m; row++ {
		curr := make([]int, 1, row)
		curr[0] = 1

		mid := (row / 2) + 1
		var right int
		if math.Mod(float64(row), 2.0) == 0 {
			right = mid - 1
		} else {
			right = mid
		}

		for j := 1; j < mid; j++ {
			curr = append(curr, prev[j-1]+prev[j])
		}

		s := make([]int, right)
		r := curr[0:right]
		copy(s, r)
		rev := reverse(s)
		curr = append(curr, rev...)
		fmt.Println(curr)
		prev = curr
	}
}
