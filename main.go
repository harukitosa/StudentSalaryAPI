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
	dns := "root:@tcp(127.0.0.1:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
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

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				Review:   review,
				Workdata: workdata,
				Company:  company,
			}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == "http://localhost:3000"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	log.Println("server start :8080")

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}

// e := echo.New()

// e.Use(middleware.Logger())
// e.Use(middleware.Recover())
// localhostの方は消す

// var db *sqlx.DB

// if err != nil {
// 	log.Fatal(err)
// }

// reviewHandler := handler.NewReviewHandler(review)
// workdataHandler := handler.NewWorkDataHandler(workdata)

// e.GET("/", func(c echo.Context) error {
// 	return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
// })

// // // JobSalary
// e.GET("/jobSalary", workdataHandler.GetReview)
// e.GET("/jobSalary/statistics", workdataHandler.GetAggregateWorkData)
// e.POST("/jobSalary", workdataHandler.CreateWorkData)
// // e.POST("/jobSalaries", jobSalaryAPI.ExportJobsSalary)

// // Review
// e.GET("/review", reviewHandler.GetReview)
// e.GET("/review/:id", reviewHandler.GetReviewByID)
// e.POST("/review", reviewHandler.CreateReview)
// e.GET("/review/created", reviewHandler.GetReviewByCreated)
// e.POST("/reviews", reviewAPI.ExportReview)

// // JobSalaryMap
// e.GET("/jobSalaryMap", jobSalaryMapAPI.GetJobSalaryMap)
// e.GET("/jobSalaryMap/count", jobSalaryMapAPI.GetJobSalaryMapByCount)
// e.GET("/jobSalaryMapData", jobSalaryAPI.GetJobsSalaryMap)

// port := os.Getenv("PORT")
// if port == "" {
// 	port = "8080"
// 	log.Printf("Defaulting to port %s", port)
// }
// log.Printf("Listening on port %s", port)
// e.Logger.Fatal(e.Start(":" + port))
