package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestGetRedigo(t *testing.T) {
	t.Run("miniRedis", func(t *testing.T) {
		conn := GetRedigo()
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
		post := GinPost("/test", test, A)
		fmt.Println(string(post))
	})
}

func B(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func TestGinGet(t *testing.T) {
	t.Run("ginGet", func(t *testing.T) {
		get := GinGet("/test", B)
		fmt.Println(string(get))
	})
}
