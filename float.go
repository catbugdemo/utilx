package main

import "math"

const MIN = 0.000001

// FloatIsEqual
// 判断 float 是否相等
func FloatIsEqual(x, y float64) bool {
	return math.Abs(x-y) < MIN
}
