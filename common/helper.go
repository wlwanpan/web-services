package helper

import (
	"log"
	"os"
	"strconv"
)

func GetEnv(k string, d string) string {
	if value, ok := os.LookupEnv(k); ok {
		return value
	}
	return d
}

func StrToUint(s string) uint64 {
	parsedS, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Println("Error Parsing to Uint: ", s)
		return 0
	}
	return parsedS
}

func StrToInt(s string) int64 {
	parsedS, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("Error Parsing to Int: ", s)
		return 0
	}
	return parsedS
}
