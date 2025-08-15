package cache

import "strings"

// GenerateKey
//
// Generate redis key.
func GenerateKey(service string, module string, keys ...string) string {
	return service + ":" + module + ":" + strings.Join(keys, ":")
}
