package main

import "testing"

func TestMain(t *testing.T) {
	word := Say()
	if word != "vim-go" {
		t.Error("not implemented")
	}
}
