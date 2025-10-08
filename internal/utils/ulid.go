package internalutils

import (
	"regexp"
)

// cUlidRegex contains the regex for valid ULID
const cUlidRegex = "^[0-7][0-9A-HJKMNP-TV-Z]{25}$"

// IsWellFormedUlidString returns whethr the ulidString is a properly formatted ulid string
func IsWellFormedUlidString(ulidString string) bool {
	re := regexp.MustCompile(cUlidRegex)
	return re.MatchString(ulidString)
}
