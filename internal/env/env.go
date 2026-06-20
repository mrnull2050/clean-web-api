package env

import (
	"os"
	"strconv"
)

func GetEnvString(Key, DefaultValue string) string {
	if value, exites := os.LookupEnv(Key); exites {
		return value
	}
	return DefaultValue
}
func GetEnvInt(Key string, DefaultValue int) int {
	if value, exites := os.LookupEnv(Key); exites {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue

		}
	}
	return DefaultValue
}
