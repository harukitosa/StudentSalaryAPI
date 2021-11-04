package main

// // [START import]
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"studentSalaryAPI/graph"
	"studentSalaryAPI/graph/generated"
	"studentSalaryAPI/infra"
	"time"

	"cloud.google.com/go/datastore"
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

var datastoreClient *datastore.Client

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
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080", "https://studentsalary.vercel.app", "https://student-salary.com", "https://www.student-salary.com"},
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
				return r.Host == "http://localhost:3000" || r.Host == "https://student-salary.com" || r.Host == "https://studentsalary.vercel.app"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	ctx := context.Background()

	datastoreClient, err = datastore.NewClient(ctx, os.Getenv("GCLOUD_DATASET_ID"))
	if err != nil {
		log.Fatal(err)
	}

	router.Handle("/", playground.Handler("student-salary", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/handle", handle)

	log.Println("server start :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Get a list of the most recent visits.
	visits, err := queryVisits(ctx, 10)
	if err != nil {
		msg := fmt.Sprintf("Could not save visit: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Record this visit.
	if err := recordVisit(ctx, time.Now(), r.RemoteAddr); err != nil {
		msg := fmt.Sprintf("Could not save visit: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Previous visits:")
	for _, v := range visits {
		fmt.Fprintf(w, "[%s] %s\n", v.Timestamp, v.UserIP)
	}

	fmt.Fprintln(w, "\nSuccessfully stored an entry of the current request.")
}

type visit struct {
	Timestamp time.Time
	UserIP    string
}

func recordVisit(ctx context.Context, now time.Time, userIP string) error {
	v := &visit{
		Timestamp: now,
		UserIP:    userIP,
	}

	k := datastore.IncompleteKey("Visit", nil)

	_, err := datastoreClient.Put(ctx, k, v)
	return err
}

func queryVisits(ctx context.Context, limit int64) ([]*visit, error) {
	q := datastore.NewQuery("Visit").Order("-Timestamp").Limit(10)

	visits := make([]*visit, 0)
	_, err := datastoreClient.GetAll(ctx, q, &visits)
	return visits, err
}
