package longestcommonprefix14

import "strings"

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	var builder strings.Builder
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || first[i] != strs[j][i] {
				return builder.String()
			}
		}
		builder.WriteByte(first[i])
	}
	return builder.String()
}
