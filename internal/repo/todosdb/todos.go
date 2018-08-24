package todosdb

import (
	"context"
	"time"

	"github.com/ges-sh/dbug/dbugdb"
)

type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func CreateTodo(ctx context.Context, db dbugdb.DB, name string, description string) (Todo, error) {
	var t Todo
	err := db.QueryRowContext(ctx, `
	INSERT INTO todos(name, description) 
			VALUES($1,$2) 
			RETURNING id,name,description, created_at,updated_at
			`, name, description).Scan(
		&t.ID,
		&t.Name,
		&t.Description,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	return t, err
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanTodo(s scanner) (Todo, error) {
	var t Todo
	err := s.Scan(
		&t.ID,
		&t.Name,
		&t.Description,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	return t, err
}

func FetchTodo(ctx context.Context, db dbugdb.DB, id int) (Todo, error) {
	row := db.QueryRowContext(ctx, `
	SELECT 
			id, 
			name,
			description,
	    created_at,	
			updated_at
		FROM todos 
		where id = $1
		`, id,
	)

	return scanTodo(row)
}

func FetchTodos(ctx context.Context, db dbugdb.DB) ([]Todo, error) {
	todos := []Todo{}
	rows, err := db.QueryContext(ctx, `
	SELECT 
			id, 
			name, 
			description, 
			created_at, 
			updated_at 
		FROM todos
	`)
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		todo, err := scanTodo(rows)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()
}

func DeleteTodo(ctx context.Context, db dbugdb.DB, id int) error {
	_, err := db.ExecContext(ctx, `
		DELETE 
				FROM todos 
				WHERE id = $1
			`, id)
	return err
}

func UpdateTodo(ctx context.Context, db dbugdb.DB, id int, name string, description string) error {
	_, err := db.ExecContext(ctx, `
UPDATE todos 
		SET 
			name = $1, 
			description = $2, 
			updated_at = now() 
		WHERE id = $3
		`, name, description, id,
	)
	return err
}
