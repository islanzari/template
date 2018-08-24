package todos

import (
	"context"

	"github.com/islanzari/template/internal/repo/todosdb"
)

type Todos interface {
	CreateTodo(ctx context.Context, name, description string) (todosdb.Todo, error)
	FetchTodo(ctx context.Context, id int) (todosdb.Todo, error)
	FetchTodos(ctx context.Context) ([]todosdb.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
	UpdateTodo(ctx context.Context, id int, name, description string) error
}

type Handle struct {
	Todos Todos
}
