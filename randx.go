package main

import (
	"math/rand"
	"time"
)

// 生成一个 r 结构体的伪随机数
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Num
// 生成一个随机数 ，范围 [min,max)
func Num(min, max int) int {
	if min > max {
		return Num(max, min)
	}
	return r.Intn(max-min) + min
}
