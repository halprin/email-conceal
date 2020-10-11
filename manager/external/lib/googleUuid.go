package lib

import (
	"github.com/google/uuid"
)

type GoogleUuid struct {}

func (receiver GoogleUuid) GenerateRandomUuid() string {
	randomUuid := uuid.New()
	return randomUuid.String()
}
