package crt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Split splits a string on the first occurrence of sep, replacing wrangler's kv.Split
func Split(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// RSplit splits a string on the last occurrence of sep, replacing wrangler's kv.RSplit
func RSplit(s, sep string) (string, string) {
	i := strings.LastIndex(s, sep)
	if i < 0 {
		return "", s
	}
	return s[:i], s[i+len(sep):]
}

// ContainsString checks if a string slice contains a given string, replacing wrangler's slice.ContainsString
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// SafeConcatName concatenates name parts with hashing to ensure the result
// fits within Kubernetes name length limits, replacing wrangler's name.SafeConcatName
func SafeConcatName(parts ...string) string {
	fullName := strings.Join(parts, "-")
	if len(fullName) <= 63 {
		return fullName
	}
	hash := sha256.Sum256([]byte(fullName))
	hashStr := hex.EncodeToString(hash[:])[:8]
	truncated := fullName[:63-9]
	return fmt.Sprintf("%s-%s", truncated, hashStr)
}

// Limit truncates a name to maxLen, replacing wrangler's name.Limit
func Limit(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
