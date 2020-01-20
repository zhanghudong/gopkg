package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	return strings.ReplaceAll(uuid.Must(uuid.NewV4()).String(), "-", "")
}
