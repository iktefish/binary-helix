package main

import (
	"testing"

	"github.com/iktefish/binary-helix/analyser"
)

func TestExactMatch(t *testing.T) {
	in := "ACCATGGATGA"
	p := "ACCA"

	out := analyser.ExactMatch(p, in)
	if out[0] != 0 {
		t.Error("Error")
	}

	p = "TG"
	out = analyser.ExactMatch(p, in)
	if out[0] != 4 {
		t.Error("Error")
	}

	p = "GGA"
	out = analyser.ExactMatch(p, in)
	if out[0] != 5 {
		t.Error("Error")
	}
}

func BenchmarkExactMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := "ACCATGGATGA"
		p := "ACCA"

		analyser.ExactMatch(p, in)
	}
}
