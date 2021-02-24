package main

import (
	_ "fmt"
	"testing"
)

func TestBasicReturn(t *testing.T) {
	ret, ret2 := returnWithNames()
	if ret != "first" {
		t.Fatalf("failed %s, %s", ret, ret2)
	}
}
