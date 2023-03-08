package data

import (
	"time"
)

type school struct {
	ID        int64
	Name      string
	Level     string
	Contact   string
	Phone     string
	Email     string
	Website   string
	Address   string
	Mode      []string
	CreatedAt time.Time
	Version   int32
}
