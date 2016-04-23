package main

import (
	"testing"
)

func TestInitAction(t *testing.T) {
	t.Error("no create test")
}
func TestNewAction(t *testing.T) {
	t.Error("no create test")
}
func TestCreateFilename(t *testing.T) {
	if filename := createFilename("test"); filename != "test.txt" {
		t.Error("func result failed")
	}
}
