migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/resti?sslmode=disable" -verbose up 

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/resti?sslmode=disable" -verbose down


run:
	go run cmd/main.go