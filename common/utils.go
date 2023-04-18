package common

import "github.com/google/uuid"

func GenerateToken() string {
	code := uuid.New().String()
	return code
}
