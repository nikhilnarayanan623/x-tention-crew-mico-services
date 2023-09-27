package response

import "time"

type User struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AllUsers struct {
	Count uint64
	Names []string
}
