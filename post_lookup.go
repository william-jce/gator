package main

import (
	"github.com/google/uuid"
)

const shortIDLen = 8

func shortID(id uuid.UUID) string {
	return id.String()[:shortIDLen]
}
