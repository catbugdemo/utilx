package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// HttpPostJSON
// url 地址 ，value 需要传递的数据 ，  dest 需要绑定的数据
func HttpPostJSON(url string, value interface{}, dest interface{}) error {
	buf, e := json.Marshal(value)
	if e != nil {
		return e
	}
	resp, e := http.Post(url, "application/json", bytes.NewBuffer(buf))
	defer resp.Body.Close()
	if e != nil {
		return e
	}
	if resp == nil {
		return errors.New("resp is nil")
	}
	if resp.Body == nil {
		return errors.New("resp.Body is nil")
	}

	if e = json.NewDecoder(resp.Body).Decode(dest); e != nil {
		return e
	}
	return nil
}
