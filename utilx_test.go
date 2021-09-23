package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"path"
	"testing"
)

func TestFile(t *testing.T) {
	getenv := os.Getenv(`GOPATH`)
	
	t.Run("CheckNotExist", func(t *testing.T) {
		exist := CheckNotExist(path.Join(getenv, `src/github.com/catbugdemo`, `file.go`))
		assert.False(t, exist)
	})
	
	t.Run(`Mkdir`, func(t *testing.T) {
		Mkdir(path.Join(getenv,))
	})
}


func TestFloat(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		equal := FloatIsEqual(1.2, 1.21)
		assert.False(t, equal)
	})

	t.Run("true", func(t *testing.T) {
		equal := FloatIsEqual(1.201, 1.201)
		assert.True(t, equal)
	})
}

func TestJsonx(t *testing.T) {
	type Test struct {
		Name string
	}
	t.Run("JsonFormatStr", func(t *testing.T) {
		test := Test{
			Name: "123",
		}
		json, e := JsonFormatStr(test)
		assert.Nil(t, e)
		fmt.Println(json)
	})
}

func TestRandx(t *testing.T) {
	t.Run("min < max", func(t *testing.T) {
		num := RandNum(1, 10)
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, num)
	})
	t.Run("min > max", func(t *testing.T) {
		num := RandNum(10, 1)
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, num)
	})
}


func TestGetRedigo(t *testing.T) {
	t.Run("miniRedis", func(t *testing.T) {
		conn := TGetRedigo()
		assert.NotNil(t, conn)
	})
}

func A(c *gin.Context) {
	type Param struct {
		Name string
	}
	var param Param
	if e := c.ShouldBindJSON(&param); e != nil {
		log.Println("fail")
	}
	c.JSON(http.StatusOK, gin.H{"Name": param.Name})
}

func TestGinPost(t *testing.T) {

	t.Run("ginPost", func(t *testing.T) {
		type Test struct {
			Name string
		}
		test := Test{Name: "123"}
		post := TGinPost("/test", test, A)
		fmt.Println(string(post))
	})
}

func B(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func TestGinGet(t *testing.T) {
	t.Run("ginGet", func(t *testing.T) {
		get := TGinGet("/test", B)
		fmt.Println(string(get))
	})
}






