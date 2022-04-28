package analyser

func ReverseComplement(s string) string {
	complement := make(map[string]string)

	complement["A"] = "T"
	complement["C"] = "G"
	complement["G"] = "C"
	complement["T"] = "A"
	complement["N"] = "N"

	rc := ""

	for _, base := range s {
		rc = complement[string(base)] + rc
	}

	return rc
}
