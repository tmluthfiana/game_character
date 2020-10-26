package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "name") {
		return errors.New("Name Already Taken")
	}

	if strings.Contains(err, "character_code") {
		return errors.New("Incorrect character_code")
	}

	return errors.New("Incorrect Details")
}
