package main

import (
	"testing"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/types"
)

func TestBoyerMoore(t *testing.T) {
	in := "GCTACGATCTAGAATCTA"
	pattern := "TCTA"
	pBM := types.ConstructBM(pattern)

	out := analyser.BoyerMoore(pattern, pBM, in)

	for _, i := range out {
		if string(in[i]) != "T" {
			t.Error("`out` must be [7, 14]\nExpected T")
		}
	}

	i := len(out) - 1
	if out[i] != 14 {
		t.Error("`out` must be [7, 14]\nExpected 14")
	}
}

func BenchmarkBoyerMoore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := "GCTACGATCTAGAATCTA"
		pattern := "TCTA"
		pBM := types.ConstructBM(pattern)

		analyser.BoyerMoore(pattern, pBM, in)
	}
}
