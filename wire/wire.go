package wire

import (
	"studentSalaryAPI/application"
	"studentSalaryAPI/handler"
	"studentSalaryAPI/infra/mysql"

	"github.com/jmoiron/sqlx"
)

// InitUserAPI is
func InitUserAPI(db *sqlx.DB) handler.UserHandler {
	userRepository := mysql.NewUserRepositoryImpl(db)
	userApplication := application.NewUserApplication(userRepository)
	userHandler := handler.NewUserHandler(userApplication)
	return userHandler
}
