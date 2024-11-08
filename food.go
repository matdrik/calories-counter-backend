// TODO - перенести в ./internal

package calories_counter_backend

type FoodResponse struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Carbs    string `json:"carbs" db:"carbs"`
	Calories string `json:"calories" db:"calories"`
	Fat      string `json:"fat" db:"fat"`
	Protein  string `json:"protein" db:"protein"`
}
