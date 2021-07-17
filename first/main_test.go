// Copyright 2021 Xage Security, Inc. All rights reserved.

package main

import (
	"fmt"
	"testing"
)

func Composition(t *testing.T) {
	t.Run("Composition Test", func(t *testing.T) {

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
