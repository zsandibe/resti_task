migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/resti?sslmode=disable" -verbose up 

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/resti?sslmode=disable" -verbose down

.PHONY: server
run:
	go run cmd/main.go

.PHONY: docker-up
docker-up:
	docker compose up

.PHONY: docker-down
docker-down:
	docker compose down



.PHONY: deps server
swag_ui:
	@echo "Open swagger index.html"
	open http://localhost:8888/swagger/index.html