package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	calories_counter_backend "server"
)

type FoodPostgres struct {
	db *pgx.Conn
}

func NewFoodPostgres(db *pgx.Conn) *FoodPostgres {
	return &FoodPostgres{db: db}
}

func (r *FoodPostgres) GetAll() ([]calories_counter_backend.FoodResponse, error) {
	var data = make([]calories_counter_backend.FoodResponse, 0)

	query := fmt.Sprintf("SELECT ft.id, ft.name, ft.calories, ft.protein, ft.fat, ft.carbs FROM %s ft", FoodTable)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item calories_counter_backend.FoodResponse
		err = rows.Scan(
			&item.ID, &item.Name, &item.Calories, &item.Protein, &item.Fat, &item.Carbs)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, err
}
