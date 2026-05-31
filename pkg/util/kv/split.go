package kv

import "strings"

// RSplit splits s by sep. If there is only one part, returns ("", part).
func RSplit(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	if len(parts) == 1 {
		return "", strings.TrimSpace(parts[0])
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(safeIndex(parts, 1))
}

// Split splits s by sep and returns the two parts.
func Split(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	return strings.TrimSpace(parts[0]), strings.TrimSpace(safeIndex(parts, 1))
}

// SplitLast splits s at the last occurrence of sep.
func SplitLast(s, sep string) (string, string) {
	idx := strings.LastIndex(s, sep)
	if idx > -1 {
		return strings.TrimSpace(s[:idx]), strings.TrimSpace(s[idx+1:])
	}
	return s, ""
}

// SplitMap splits s by sep into key=value pairs.
func SplitMap(s, sep string) map[string]string {
	return SplitMapFromSlice(strings.Split(s, sep))
}

// SplitMapFromSlice builds a map from "key=value" strings.
func SplitMapFromSlice(parts []string) map[string]string {
	result := map[string]string{}
	for _, part := range parts {
		k, v := Split(part, "=")
		result[k] = v
	}
	return result
}

func safeIndex(parts []string, idx int) string {
	if len(parts) <= idx {
		return ""
	}
	return parts[idx]
}
