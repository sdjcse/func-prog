package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_composition(t *testing.T) {
	t.Run("Composition Test", func(t *testing.T) {
		assert.Equal(t, "4.500000", composition(numberFunc, stringFun)(3))
		assert.Equal(t, "1.500000", composition(numberFunc, stringFun)(1))
		assert.Equal(t, "0.000000", composition(numberFunc, stringFun)(0))
		assert.Equal(t, "3.000000", composition(numberFunc, stringFun)(2))
		assert.Equal(t, "-15.000000", composition(numberFunc, stringFun)(-10))
	})
}

func numberFunc(test int) float32 {
	return float32(test) * float32(1.5)
}

func stringFun(number float32) string {
	return fmt.Sprintf("%f", number)
}

func Test_id(t *testing.T) {
	t.Run("idempotency test", func(t *testing.T) {
		assert.Equal(t, "abc", id("abc").(string))
		assert.NotEqual(t, "def", id(123).(int))
	})
}
