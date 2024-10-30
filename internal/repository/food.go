package repository

import (
	"context"
	"fmt"
	"server/internal/models"
)

func (repo *PGRepo) GetFood() ([]models.Food, error) {
	rows, err := repo.pool.Query(context.Background(), "SELECT id, name, calories, fat, carbs, protein FROM food")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Food
	for rows.Next() {
		var item models.Food
		err = rows.Scan(
			&item.ID, &item.Name, &item.Calories, &item.Fat, &item.Carbs, &item.Protein)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, err
}

func (repo *PGRepo) CreateFood(newFood models.Food) error {
	const queryString = `INSERT INTO food(name, calories, fat, carbs, protein) VALUES ($1, $2, $3, $4, $5)`
	_, err := repo.pool.Exec(context.Background(), queryString, newFood.Name, newFood.Calories, newFood.Fat, newFood.Carbs, newFood.Protein)
	if err != nil {
		return fmt.Errorf("не удалось добавить продукт: %w", err)
	}
	return nil
}
