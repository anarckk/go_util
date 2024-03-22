package go_uuid

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	u2 := uuid.NewV4()
	return strings.ReplaceAll(u2.String(), "-", "")
}
