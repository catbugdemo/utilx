package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
