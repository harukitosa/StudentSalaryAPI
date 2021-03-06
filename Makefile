test:
	go test ./...

build:
	go build

deploy:
	gcloud app deploy

dev:
	go build && ./studentSalaryAPI

db:
	./cloud_sql_proxy -instances=student-salary-api:asia-northeast1:student-salary-api=tcp:3306