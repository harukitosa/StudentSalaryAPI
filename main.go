package main

// // [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studentSalaryAPI/domain"
	"studentSalaryAPI/wire"

	_ "github.com/go-sql-driver/mysql"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func initDB() *gorm.DB {
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dns := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initLocalDB() *gorm.DB {

	// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	// if !isSet {
	// 	socketDir = "/cloudsql"
	// }

	dns := "root:@tcp(127.0.0.1:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// localhostの方は消す

	v := os.Getenv("RUNENV")

	var db *gorm.DB
	if v == "production" {
		db = initDB()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://student-salary.com"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	} else {
		db = initLocalDB()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	}

	userAPI := wire.InitUserAPI(db)
	jobSalaryAPI := wire.InitJobSalaryAPI(db)
	jobSalaryMapAPI := wire.InitJobSalaryMapAPI(db)
	reviewAPI := wire.InitReviewAPI(db)

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.JobSalary{})
	db.AutoMigrate(&domain.Review{})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})
	e.GET("/data", userAPI.CreateUser)
	e.GET("/get", userAPI.GetAllUser)

	// JobSalary
	e.GET("/jobSalary", jobSalaryAPI.GetAllJobSalary)
	e.GET("/jobSalary/statistics", jobSalaryAPI.GetStatistics)
	e.POST("/jobSalary", jobSalaryAPI.CreateJobSalary)
	e.POST("/jobSalaries", jobSalaryAPI.ExportJobsSalary)

	// Review
	e.GET("/review", reviewAPI.GetAllReview)
	e.GET("/review/:id", reviewAPI.GetReviewByID)
	e.GET("/review/created", reviewAPI.GetReviewByCreated)
	e.POST("/review", reviewAPI.CreateReview)
	e.POST("/reviews", reviewAPI.ExportReview)

	// JobSalaryMap
	e.GET("/jobSalaryMap", jobSalaryMapAPI.GetJobSalaryMap)
	e.GET("/jobSalaryMap/count", jobSalaryMapAPI.GetJobSalaryMapByCount)

	e.GET("/jobSalaryMapData", jobSalaryAPI.GetJobsSalaryMap)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
