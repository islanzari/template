package controller

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/islanzari/template/internal/config"
	"github.com/islanzari/template/internal/controller/middleware"
	"github.com/islanzari/template/internal/controller/todos"
	"github.com/islanzari/template/internal/controller/users"
	"github.com/islanzari/template/internal/model"
	"github.com/julienschmidt/httprouter"
	"github.com/lhecker/argon2"
	"github.com/sirupsen/logrus"
)

// Controller is an main router of application.
func Controller(ctx context.Context, db *sql.DB, config config.Config, log logrus.FieldLogger) http.Handler {
	router := httprouter.New()

	hasher := argon2.DefaultConfig()

	usersModel := model.Users{
		DB:     db,
		Hasher: &hasher,
	}
	todoModel := model.Todos{
		DB: db,
	}
	users := users.Handle{
		Users:  usersModel,
		Config: config,
	}
	todos := todos.Handle{
		Todos: todoModel,
	}
	router.POST("/api/todos/", todos.CreateTodo)
	router.GET("/api/todos/:id/", todos.FetchTodo)
	router.GET("/api/todos/", todos.FetchTodos)
	router.PATCH("/api/todos/:id/", todos.UpdateTodo)
	router.DELETE("/api/todos/:id/", todos.DeleteTodo)

	router.POST("/api/users/", users.Create)
	router.GET("/api/me/", users.Me)
	router.POST("/api/session/", users.Login)

	loggerMiddleware := middleware.Logger{
		Log:  log,
		Next: router,
	}

	authenticateMiddleware := middleware.Authenticate{
		Users: usersModel,
		Next:  loggerMiddleware,
	}

	jwtMiddleware := middleware.JWT{
		JWTSecret: config.JWTSecret,
		Next:      authenticateMiddleware,
	}

	return jwtMiddleware
}
