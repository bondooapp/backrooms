package cache

import "strings"

func GenerateKey(service string, module string, keys ...string) string {
	return service + ":" + module + ":" + strings.Join(keys, ":")
}
