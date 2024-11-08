// TODO - перенести в ./internal

package calories_counter_backend

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
