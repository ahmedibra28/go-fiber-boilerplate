package utils

import (
	"time"

	"github.com/google/uuid"
)

func GenerateResetToken() (string, time.Time) {
	expires := time.Now().Add(10 * time.Minute)
	token := uuid.New().String()
	return token, expires
}
