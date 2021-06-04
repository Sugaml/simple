package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "invalid character") {
		return errors.New("Invalid Character")
	}
	if strings.Contains(err, "json") {
		return errors.New("Json Error")
	}
	return errors.New("Incorrect Details")
}
