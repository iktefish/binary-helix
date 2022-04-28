package main

import (
	"testing"

	"github.com/iktefish/binary-helix/analyser"
)

func TestComplement(t *testing.T) {
	if analyser.Complement("A") != "T" {
		t.Error("Expected T")
	}
	if analyser.Complement("T") != "A" {
		t.Error("Expected A")
	}
	if analyser.Complement("C") != "G" {
		t.Error("Expected G")
	}
	if analyser.Complement("G") != "C" {
		t.Error("Expected C")
	}
	if analyser.Complement("N") != "N" {
		t.Error("Expected N")
	}
}

func BenchmarkComplement(b *testing.B) {
	list := []string{"A", "T", "C", "G", "N"}
	for i := 0; i < b.N; i++ {
		for _, r := range list {
			analyser.Complement(r)
		}
	}
}
