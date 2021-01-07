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

// InitJobSalaryAPI is
func InitJobSalaryAPI(db *gorm.DB) handler.JobSalaryHandler {
	jobSalaryRepository := mysql.NewJobSalaryRepositoryImpl(db)
	jobSalaryApplication := application.NewJobSalaryApplication(jobSalaryRepository)
	jobSalaryHandler := handler.NewJobSalaryHandler(jobSalaryApplication)
	return jobSalaryHandler
}

// InitJobSalaryMapAPI is
func InitJobSalaryMapAPI(db *gorm.DB) handler.JobSalaryMapHandler {
	jobSalaryMapRepository := mysql.NewJobSalaryMapRepositoryImpl(db)
	jobSalaryMapApplication := application.NewJobSalaryMapApplication(jobSalaryMapRepository)
	jobSalaryMapHandler := handler.NewJobSalaryMapHandler(jobSalaryMapApplication)
	return jobSalaryMapHandler
}
