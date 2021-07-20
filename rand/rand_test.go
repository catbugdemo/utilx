package rand

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNum(t *testing.T) {
	t.Run("min < max", func(t *testing.T) {
		num := Num(1, 10)
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, num)
	})
	t.Run("min > max", func(t *testing.T) {
		num := Num(10, 1)
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, num)
	})
}
