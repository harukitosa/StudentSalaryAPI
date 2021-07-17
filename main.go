package main

// // [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studentSalaryAPI/handler"
	"studentSalaryAPI/infra"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func initDB() (*sqlx.DB, error) {
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
	db, err := sqlx.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initLocalDB() (*sqlx.DB, error) {
	dns := "root:@tcp(127.0.0.1:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// localhostの方は消す

	v := os.Getenv("RUNENV")

	var db *sqlx.DB
	var err error
	if v == "production" {
		db, err = initDB()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://student-salary.com"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	} else {
		db, err = initLocalDB()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	}

	if err != nil {
		log.Fatal(err)
	}

	review := infra.NewReviewInfra(db)
	workdata := infra.NewWorkDataInfra(db)
	reviewHandler := handler.NewReviewHandler(review)
	workdataHandler := handler.NewWorkDataHandler(workdata)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})

	// // JobSalary
	e.GET("/jobSalary", workdataHandler.GetReview)
	// e.GET("/jobSalary/statistics", jobSalaryAPI.GetStatistics)
	e.POST("/jobSalary", workdataHandler.CreateWorkData)
	// e.POST("/jobSalaries", jobSalaryAPI.ExportJobsSalary)

	// Review
	e.GET("/review", reviewHandler.GetReview)
	e.GET("/review/:id", reviewHandler.GetReviewByID)
	e.POST("/review", reviewHandler.CreateReview)
	// e.GET("/review/created", reviewAPI.GetReviewByCreated)
	// e.POST("/review", reviewAPI.CreateReview)
	// e.POST("/reviews", reviewAPI.ExportReview)

	// // JobSalaryMap
	// e.GET("/jobSalaryMap", jobSalaryMapAPI.GetJobSalaryMap)
	// e.GET("/jobSalaryMap/count", jobSalaryMapAPI.GetJobSalaryMapByCount)

	// e.GET("/jobSalaryMapData", jobSalaryAPI.GetJobsSalaryMap)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
