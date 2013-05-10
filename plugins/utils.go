package plugins

import "strings"

func removeExtension(s string) string {
	if idx := strings.LastIndex(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
