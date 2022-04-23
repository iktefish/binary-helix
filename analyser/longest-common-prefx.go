package analyser

func LongestCommonPrefix(s1 string, s2 string) string {
	commonPrefix := 0
	for i := 0; i < len(s1) && i < len(s2) && s1[i] == s2[i]; i++ {
		commonPrefix += 1
	}
	return s1[:commonPrefix]
}
