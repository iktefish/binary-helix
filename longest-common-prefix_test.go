package main

import (
	"testing"

	"github.com/iktefish/binary-helix/analyser"
)

func TestLongestCommonPrefix(t *testing.T) {
	in := "ACCATGGATGA"
	p := "ACCATGGATGA"

	out := analyser.LongestCommonPrefix(p, in)
	if out != "ACCATGGATGA" {
		t.Error("Error")
	}

	p = "ACGCAAGTTGA"

	out = analyser.LongestCommonPrefix(p, in)
	if out == "ACGCAAGTTGA" {
		t.Error("Error")
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := "ACCATGGATGA"
		p := "ACCATGGATGA"
		analyser.LongestCommonPrefix(p, in)

		p = "ACGCAAGTTGA"
		analyser.LongestCommonPrefix(p, in)
	}
}
