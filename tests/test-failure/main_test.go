package main

import "testing"

func TestFailure(t *testing.T) {
	t.Errorf("expected failure")
}
