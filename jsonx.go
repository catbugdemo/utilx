package main

import (
	"encoding/json"
)

// JsonFormatStr 以 jsonx 格式化返回结构体
// v- 结构体
func JsonFormatStr(v interface{}) (string, error) {
	indent, e := json.MarshalIndent(v, "", "\t")
	if e != nil {
		return "", e
	}
	return string(indent), nil
}
