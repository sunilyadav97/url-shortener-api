package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"time"
)

// GenerateShortURL generates a short hash for the given URL using SHA1 and base64 encoding.
func GenerateShortURL(url string) string {
	h := sha1.New()
	h.Write([]byte(url + time.Now().String()))
	sha1Hash := h.Sum(nil)
	// Use base64 URL encoding and take the first 7 characters
	encoded := base64.URLEncoding.EncodeToString(sha1Hash)
	encoded = strings.TrimRight(encoded, "=")
	if len(encoded) > 7 {
		return encoded[:7]
	}
	return encoded
}
