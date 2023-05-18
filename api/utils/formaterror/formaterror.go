package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Alreadey Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Alreadey Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Alreadey Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}

	return errors.New("Incorrect Details")
}
