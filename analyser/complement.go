package analyser

func Complement(s string) string {
	complement := make(map[string]string)

	complement["A"] = "T"
	complement["C"] = "G"
	complement["G"] = "C"
	complement["T"] = "A"
	complement["N"] = "N"

	c := ""

	for _, base := range s {
		c = c + complement[string(base)]
	}

	return c
}
