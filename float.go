package main

import "math"

// MIN 最小判断精度
const MIN = 0.000001

// FloatIsEqual  判断 float 是否相等
func FloatIsEqual(x, y float64) bool {
	return math.Abs(x-y) < MIN
}
