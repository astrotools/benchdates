package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	num = flag.Int("n", 10000, "Size of test dates.")
)

func main() {
	flag.Parse()

	lower := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	upper := time.Now()
	lowerUnix := lower.Unix()
	upperUnix := upper.Unix()

	seed := upper.UnixNano()
	rng := rand.New(rand.NewSource(seed))
	n := abs(lowerUnix) + upperUnix

	for i := 0; i < *num; i++ {
		secs := rng.Int63n(n) + lowerUnix
		fmt.Println(secs)
	}
}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}
