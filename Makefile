test:
	go test ./...

build:
	go build

deploy:
	gcloud app deploy --project=studentsalaryprod

gql:
	go run github.com/99designs/gqlgen generate


dev:
	go build && ./studentSalaryAPI

db:
	./cloud_sql_proxy -instances=student-salary-api:asia-northeast1:student-salary-api=tcp:3306

migrate-up:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/sample' -path db/migrations up

migrate-down:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/sample' -path db/migrations down 

migrate-up-production:
	migrate -database 'mysql://student:salary@tcp(127.0.0.1:3306)/student_salary_production' -path db/migrations up

migrate-down-production:
	migrate -database 'mysql://student:salary@tcp(127.0.0.1:3306)/student_salary_production' -path db/migrations down