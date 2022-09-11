package utils

import "github.com/satori/go.uuid"

func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
