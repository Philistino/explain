package main

import (
	"fmt"
	"testing"
)

func TestAnsiCounter(t *testing.T) {
	t.Parallel()
	counter := ansiRepeater()
	// 18 is an arbitrary number greater than 14
	colors := make([]int, 18)
	for i := 1; i < 18; i++ {
		ref := fmt.Sprintf("helpRef-%d", i)
		color := counter(ref)
		colors[i] = color
	}
	for i := 1; i < 18; i++ {
		if i <= 14 && i != colors[i] {
			t.Error()
		}
		if i > 14 && colors[i] != i%14 {
			t.Error()
		}
	}
	color := counter("helpRef-2")
	if color != 2 {
		t.Error("Color should be 2")
	}
	color = counter("helpRef-15")
	if color != 1 {
		t.Error("Color should be 1")
	}
	color = counter("")
	if color != 15 {
		t.Error("Color should be 15")
	}
}
