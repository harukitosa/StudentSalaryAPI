package wire

import (
	"studentSalaryAPI/application"
	"studentSalaryAPI/handler"
	"studentSalaryAPI/infra/mysql"

	"gorm.io/gorm"
)

// InitUserAPI is
func InitUserAPI(db *gorm.DB) handler.UserHandler {
	userRepository := mysql.NewUserRepositoryImpl(db)
	userApplication := application.NewUserApplication(userRepository)
	userHandler := handler.NewUserHandler(userApplication)
	return userHandler
}
