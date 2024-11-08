package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	calories_counter_backend "server"
)

type FoodPostgres struct {
	db *sqlx.DB
}

func NewFoodPostgres(db *sqlx.DB) *FoodPostgres {
	return &FoodPostgres{db: db}
}

func (r *FoodPostgres) GetAll() ([]calories_counter_backend.FoodResponse, error) {
	var list = make([]calories_counter_backend.FoodResponse, 0)

	query := fmt.Sprintf("SELECT ft.id, ft.name, ft.carbs, ft.calories, ft.fat, ft.protein FROM %s ft", FoodTable)
	err := r.db.Select(&list, query)

	return list, err
}
