package util

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	u1 := uuid.Must(uuid.NewUUID()).String()
	return strings.ReplaceAll(u1, "-", "")
}
