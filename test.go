package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// TGetRedigo 测试中使用模拟缓存
func TGetRedigo() redis.Conn {
	mock, e := miniredis.Run()
	if e != nil {
		log.Fatalln("fail to get mock err:", e)
	}
	dial, _ := redis.Dial(`tcp`, mock.Addr())
	if e != nil {
		log.Fatalln("fail to dial redis err:", e)
	}
	return dial
}

// TGinPost 当测试 c.Bind() 头没有绑定 application/json 无法获取参数
// 测试 gin Post
// url 地址，value 测试数据 ， handlers
func TGinPost(url string, value interface{}, handlers gin.HandlerFunc) []byte {
	r := initGin()
	r.POST(url, handlers)

	jsonByte, e := json.Marshal(value)
	if e != nil {
		log.Fatalln("fail to marshal json err:", e)
	}

	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonByte))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()
	//读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// TGinGet 测试并返回 gin Get
func TGinGet(url string, handlers gin.HandlerFunc) []byte {
	r := initGin()
	r.GET(url, handlers)

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()
	//读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

func initGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}
