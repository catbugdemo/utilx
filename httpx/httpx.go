package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// PostJSON
// url 地址 ，value 需要传递的数据 ，  dest 需要绑定的数据
func PostJSON(url string, value interface{}, dest interface{}) error {
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

	all, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	if resp.StatusCode != 200 {
		return errors.New("status is not 200,resp body is " + string(all))
	}
	if e = json.Unmarshal(all, &dest); e != nil {
		return e
	}
	return nil
}
