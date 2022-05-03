package main

// // [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studentSalaryAPI/graph"
	"studentSalaryAPI/graph/generated"
	"studentSalaryAPI/infra"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
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
	dns := "root:root@tcp(127.0.0.1:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	v := os.Getenv("RUNENV")

	var db *sqlx.DB
	var err error
	if v == "production" {
		db, err = initDB()
	} else {
		db, err = initLocalDB()
	}

	if err != nil {
		log.Fatal(err)
	}

	review := infra.NewReviewInfra(db)
	workdata := infra.NewWorkDataInfra(db)
	company := infra.NewCompanyInfra(db)
	blog := infra.NewBlogInfra(db)

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080", "https://studentsalary.vercel.app", "https://student-salary.com", "https://www.student-salary.com", "https://student-salary-cnhpcagh2a-an.a.run.app"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				Review:   review,
				Workdata: workdata,
				Company:  company,
				Blog:     blog,
			}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == "http://localhost:3000" || r.Host == "https://student-salary.com" || r.Host == "https://studentsalary.vercel.app" || r.Host == "https://student-salary-cnhpcagh2a-an.a.run.app"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	log.Println("server start :8080")

	router.Handle("/", playground.Handler("student-salary", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
