package models

type Food struct {
	ID       int    `db:"id" json:"-"`
	Name     string `db:"name" json:"name" binding:"required"`
	Calories int    `db:"calories" json:"calories" binding:"required"`
	Fat      int    `db:"fat" json:"fat" binding:"required"`
	Carbs    int    `db:"carbs" json:"carbs" binding:"required"`
	Protein  int    `db:"protein" json:"protein" binding:"required"`
}
