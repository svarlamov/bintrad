package utils

import "strings"

func IsRecordNotFoundError(err error) bool {
	return err.Error() == "record not found"
}

func IsRecordDuplicate(err error) bool {
	return strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry")
}
