package main

// // [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studentSalaryAPI/model"
	"studentSalaryAPI/wire"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
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

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	e := echo.New()
	v := os.Getenv("RUNENV")
	var db *gorm.DB
	if v == "production" {
		db = initDB()
	} else {
		db = initLocalDB()
	}
	userAPI := wire.InitUserAPI(db)
	jobSalaryAPI := wire.InitJobSalaryAPI(db)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.JobSalary{})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})
	e.GET("/data", userAPI.CreateUser())
	e.GET("/get", userAPI.GetAllUser())
	e.GET("/jobSalary/get", jobSalaryAPI.GetAllJobSalary())
	e.GET("/jobSalary/insert", jobSalaryAPI.CreateJobSalary())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
