package utils

import (
	"regexp"
)

var phoneNumberRegex = "^[+]?[0-9]+$"

func IsValidPhoneNumber(str string) (bool, error) {
	matched, err := regexp.MatchString(phoneNumberRegex, str)
	return matched, err
}
