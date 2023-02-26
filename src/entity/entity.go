package entity

import "time"

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	Roles     []Role    `json:"roles,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Verified  bool      `json:"-"`
}
