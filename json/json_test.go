package json

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatStr(t *testing.T) {
	type Test struct {
		Name string
	}
	t.Run("test", func(t *testing.T) {
		test := Test{
			Name: "123",
		}
		json, e := FormatStr(test)
		assert.Nil(t, e)
		fmt.Println(json)
	})
}
