package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testing_Composition(t *testing.T) {
	t.Run("Composition Test", func(t *testing.T) {
		assert.Equal(t, "4.5", composition(numberFunc, stringFun)(3))
		assert.Equal(t, "1.5", composition(numberFunc, stringFun)(1))
		assert.Equal(t, "0", composition(numberFunc, stringFun)(0))
		assert.Equal(t, "3", composition(numberFunc, stringFun)(2))
	})
}

func numberFunc(test int) float32 {
	return float32(test) * float32(1.5)
}

func stringFun(number float32) string {
	return fmt.Sprintf("%f", number)
}

func Test_id(t *testing.T) {

}
