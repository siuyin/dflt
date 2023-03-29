// Package dflt provide default values in the absence of
// defined environment variables.
package dflt

import (
	"log"
	"os"
	"strconv"
)

// EnvString sets default string dflt from a environment variable enVar.
// Eg. EnvString("MYVAR","abc") returns "abc" if MYVAR is not set or
// MYVAR's value if it is.
func EnvString(enVar, dflt string) (ret string) {
	ret = dflt
	if os.Getenv(enVar) != "" {
		ret = os.Getenv(enVar)
	}
	return
}

// EnvInt sets default int dflt from environment variable enVar.
func EnvInt(enVar string, dflt int) (ret int, err error) {
	ret = dflt
	if os.Getenv(enVar) != "" {
		ret, err = strconv.Atoi(os.Getenv(enVar))
	}
	return
}

func EnvIntMust(enVar string, dflt int) (ret int) {
	ret = dflt
	var err error

	val := os.Getenv(enVar)
	if val == "" {
		return
	}

	ret, err = strconv.Atoi(os.Getenv(enVar))
	if err != nil {
		log.Fatalf("Unable to parse %q as an integer", val)
	}
	return
}
