package dflt

import (
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
