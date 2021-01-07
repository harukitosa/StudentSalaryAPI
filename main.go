package main

// // [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studentSalaryAPI/wire"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func initDB() *sqlx.DB {
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
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initLocalDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	e := echo.New()
	v := os.Getenv("RUNENV")
	var db *sqlx.DB
	if v == "production" {
		db = initDB()
	} else {
		db = initLocalDB()
	}
	userAPI := wire.InitUserAPI(db)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})
	e.GET("/data", userAPI.CreateUser())
	e.GET("/get", userAPI.GetAllUser())

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	e.Logger.Fatal(e.Start(":" + port))
}
