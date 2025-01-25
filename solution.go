package solution

import "strings"

func solution(strs []string) string { //longestCommonPrefix
	if len(strs) == 0 {
		return ""
	}
	shortestWord := len(strs[0])
	for _, v := range strs {
		if len(v) < shortestWord {
			shortestWord = len(v)
		}
	}
	prefix := strings.Builder{}
	for i := 0; i < shortestWord; i++ {
		prefixAddition := strs[0][i]
		for _, v := range strs {
			if v[i] != prefixAddition {
				return prefix.String()
			}
		}
		prefix.WriteByte(prefixAddition)
	}
	return prefix.String()
}
