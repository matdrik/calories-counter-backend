package models

type LimitedUser struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	RoleID   int    `json:"role_id" db:"role_id"`
}

type User struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"passwordHash" db:"password_hash"`
	RoleID       int    `json:"role_id" db:"role_id"`
}
