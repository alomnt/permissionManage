package utils

import (
	"github.com/gofrs/uuid"
	"strings"
)

func GetUuid() string {
	var uuid = uuid.Must(uuid.NewV4()).String()
	uuid = strings.Replace(uuid, "-", "", -1)
	return uuid
}
