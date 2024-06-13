
all: build

build:
	@echo "Building..."
	
	@go build -o tmp/main.exe cmd/api/main.go


run:build
	@./tmp/main.exe



docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi


docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


test:
	@echo "Testing..."
	@go test ./tests -v


clean:
	@echo "Cleaning..."
	@rm -f tmp/main*


watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	       	go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi


createMigrations:
	migrate create -ext sqlx -dir db/migrations -seq $(name);\

migrateUp:
	migrate -database postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable -path db/migrations up

migrateDown:
	migrate -database postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable -path db/migrations down 

.PHONY: all build run test clean
