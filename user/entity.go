package user

import "time"

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Occupation     string    `json:"occupation"`
	PasswordHash   string    `json:"passwordHash"`
	AvatarFileName string    `json:"avatarFilename"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
