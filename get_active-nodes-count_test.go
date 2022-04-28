package main

import (
	"testing"

	"github.com/iktefish/binary-helix/utils"
)

func TestGet_ActiveNodeCount(t *testing.T) {
    if utils.Get_ActiveNodeCount() < 0 {
		t.Error("Expected >= 0")
    }
}

func BenchmarkGet_ActiveNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
        utils.Get_ActiveNodeCount()
	}
}
