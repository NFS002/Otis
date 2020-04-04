package model

import "github.com/satori/go.uuid"


/* Shared utility functions */

// GenerateUUID : Generates a random UUID.
func GenerateUUID() (uuid.UUID, error) {
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}