package json

import (
	"encoding/json"
)

// FormatStr 以 json 格式化返回结构体
// v- 结构体
func FormatStr(v interface{}) (string, error) {
	indent, e := json.MarshalIndent(v, "", "\t")
	if e != nil {
		return "", e
	}
	return string(indent), nil
}
