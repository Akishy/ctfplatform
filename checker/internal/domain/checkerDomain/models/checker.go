package models

import "github.com/google/uuid"

type Checker struct {
	UUID    uuid.UUID
	Ip      string
	WebPort int
}
