package util

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return strings.ReplaceAll(u1.String(), "-", "")
}
