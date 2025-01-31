package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := "world"
	if result != "Hello world" {
		t.Error("Expected Hello to append 'world'")
	}
}
