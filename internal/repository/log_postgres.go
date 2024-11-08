package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	calories_counter_backend "server"
	"strings"
)

type LogPostgres struct {
	db *sqlx.DB
}

func NewLogPostgres(db *sqlx.DB) *LogPostgres {
	return &LogPostgres{db: db}
}

func (r *LogPostgres) GetAll(userId int, date string) ([]calories_counter_backend.LogResponse, error) {
	var list = make([]calories_counter_backend.LogResponse, 0)

	query := fmt.Sprintf(`
		SELECT lt.id, ft.name as food, lt.meal_id, lt.quantity, lt.date, ft.calories, ft.carbs, ft.fat, ft.protein
		FROM %s lt
			INNER JOIN %s ut
		on lt.user_id = ut.id
			INNER JOIN %s ft
		on lt.food_id = ft.id
		where ut.id = $1 and lt.date = $2
	`, LogsTable, UserTable, FoodTable)

	err := r.db.Select(&list, query, userId, date)

	return list, err
}

func (r *LogPostgres) Create(userId int, log calories_counter_backend.Log) (int, error) {
	var id int
	queryString := fmt.Sprintf("insert into %s (user_id, meal_id, food_id, quantity, date) values ($1, $2, $3, $4, $5) returning id", LogsTable)
	row := r.db.QueryRow(queryString, userId, log.MealID, log.FoodID, log.Quantity, log.Date)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *LogPostgres) GetById(userId, id int) (calories_counter_backend.LogResponse, error) {
	logrus.Printf("GetById params: '%d' _ '%d'", userId, id)

	var log calories_counter_backend.LogResponse
	query := fmt.Sprintf(`
		SELECT lt.id, ft.name as food, lt.meal_id, lt.quantity, lt.date, ft.calories, ft.carbs, ft.fat, ft.protein
		FROM %s lt
			INNER JOIN %s ut
		on lt.user_id = ut.id
			INNER JOIN %s ft
		on lt.food_id = ft.id
		where lt.id = $1 and ut.id = $2
	`, LogsTable, UserTable, FoodTable)

	err := r.db.Get(&log, query, id, userId)
	return log, err
}

func (r *LogPostgres) Delete(userId, id int) error {
	queryString := fmt.Sprintf("delete from %s lt where lt.user_id = $1 and lt.id = $2", LogsTable)
	_, err := r.db.Exec(queryString, userId, id)
	return err
}

func (r *LogPostgres) Update(userId, id int, log calories_counter_backend.UpdateLogInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if log.Quantity != nil {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, *log.Quantity)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	queryString := fmt.Sprintf("update %s lt set %s from %s ut where lt.user_id = ut.id and lt.id = $%d and ut.id = $%d",
		LogsTable, setQuery, UserTable, argId, argId+1)
	args = append(args, id, userId)

	_, err := r.db.Exec(queryString, args...)
	return err
}
