package helpers

import (
	"strings"
)

func ParseConstraintError(err error) map[string]string {
	msg := err.Error()

	switch {
	case strings.Contains(msg, "email"):
		return map[string]string{
			"email": "already registered",
		}

	case strings.Contains(msg, "username"):
		return map[string]string{
			"username": "already registered",
		}

	}

	return map[string]string{
		"message": "duplicate data",
	}
}
