package controller

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/islanzari/template/internal/config"
	"github.com/islanzari/template/internal/controller/middleware"
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

	users := users.Handle{
		Users:  usersModel,
		Config: config,
	}

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
