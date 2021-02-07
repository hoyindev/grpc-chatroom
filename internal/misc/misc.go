package misc

import (
	"os"
)

//GetEnv get envirionment variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
