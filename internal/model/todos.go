package model

import (
	"context"
	"database/sql"

	"github.com/islanzari/template/internal/repo/todosdb"
)

type Todos struct {
	DB *sql.DB
}

func (t Todos) CreateTodo(ctx context.Context, name, description string) (todosdb.Todo, error) {
	return todosdb.CreateTodo(ctx, t.DB, name, description)
}

func (t Todos) FetchTodo(ctx context.Context, id int) (todosdb.Todo, error) {
	return todosdb.FetchTodo(ctx, t.DB, id)
}

func (t Todos) FetchTodos(ctx context.Context) ([]todosdb.Todo, error) {
	return todosdb.FetchTodos(ctx, t.DB)
}

func (t Todos) DeleteTodo(ctx context.Context, id int) error {
	return todosdb.DeleteTodo(ctx, t.DB, id)
}

func (t Todos) UpdateTodo(ctx context.Context, id int, name, description string) error {
	return todosdb.UpdateTodo(ctx, t.DB, id, name, description)
}
