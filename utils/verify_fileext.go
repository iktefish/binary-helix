package utils

import "path/filepath"

func Verify_FileExt(s string) bool {

	if filepath.Ext(s) == "fa" || filepath.Ext(s) == "fastq" {
		return true
	}

	return false
}
