package utils

import "path/filepath"

func Verify_FileExt(s string) bool {
	if Verify_Fasta(s) == true || Verify_Fastq(s) == true {
		return true
	}
	return false
}

func Verify_Fastq(s string) bool {
	if filepath.Ext(s) == ".fastq" {
		return true
	}
	return false
}

func Verify_Fasta(s string) bool {
	if filepath.Ext(s) == ".fa" {
		return true
	}
	return false
}
