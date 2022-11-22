package main

import (
	"fmt"
	"math"
	"sort"
)

type ExpoIdle struct {
	Diff float64
	C1   int
	C2   int
}

func main() {
	e := math.Exp(1)
	c2 := 100
	res := make([]ExpoIdle, 0, c2)
	for i := 1; i < c2; i++ {
		ifloat64 := float64(i)
		c1Floor := math.Floor(ifloat64 * e)
		c1Ceil := math.Ceil(ifloat64 * e)
		absCeil := math.Abs(c1Floor - e)
		absFloor := math.Abs(c1Ceil - e)
		if absCeil < absFloor {
			res = append(res, ExpoIdle{
				Diff: math.Abs((c1Ceil / ifloat64) - e),
				C1:   int(c1Ceil),
				C2:   int(ifloat64),
			})
		} else {
			res = append(res, ExpoIdle{
				Diff: math.Abs((c1Floor / ifloat64) - e),
				C1:   int(c1Floor),
				C2:   int(ifloat64),
			})
		}
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Diff < res[j].Diff
	})
	best := res[0]
	fmt.Println(best.C1, best.C2, best.Diff)
}
