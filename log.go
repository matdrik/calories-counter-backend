// TODO - перенести в ./internal

package calories_counter_backend

import "errors"

type LogResponse struct {
	ID       int     `json:"id" db:"id"`
	MealID   int     `json:"mealId" db:"meal_id" binding:"required"`
	Food     string  `json:"food" db:"food" binding:"required"`
	Quantity float32 `json:"quantity" db:"quantity" binding:"required"`
	Date     string  `json:"date" db:"date" binding:"required"`
	Carbs    string  `json:"carbs" db:"carbs" binding:"required"`
	Calories string  `json:"calories" db:"calories" binding:"required"`
	Fat      string  `json:"fat" db:"fat" binding:"required"`
	Protein  string  `json:"protein" db:"protein" binding:"required"`
}

type Log struct {
	ID       int     `json:"id" db:"id"`
	UserID   int     `json:"-" db:"user_id"`
	MealID   int     `json:"mealId" db:"meal_id" binding:"required"`
	FoodID   int     `json:"foodId" db:"food_id" binding:"required"`
	Quantity float32 `json:"quantity" db:"quantity" binding:"required"`
	Date     string  `json:"date" db:"date" binding:"required"`
}

type UpdateLogInput struct {
	//MealID   *int `json:"mealId" db:"meal_id"`
	//FoodID   *int `json:"foodId" db:"food_id"`
	Quantity *int `json:"quantity" db:"quantity"`
}

func (i *UpdateLogInput) Validate() error {
	if i.Quantity == nil {
		//if i.MealID == nil && i.FoodID == nil && i.Quantity == nil {
		return errors.New("поле quantity обязательно")
		//return errors.New("must provide either meal_id, food_id, quantity")
	}

	return nil
}
