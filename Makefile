include .env

# Installing frontend dependencies
.PHONY: frontend-dep
frontend-dep:
	cd frontend && npm install

# running the frontend
.PHONY: frontend-run
frontend-run:
	cd frontend && npm run dev

# Lint the frontend source code
.PHONY: frontend-lint
frontend-lint:
	cd frontend && npx eslint

# Installing backend dependencies
.PHONY: backend-dep
backend-dep:
	cd backend/cmd/server && go get .

# Lint backend source code
.PHONY: backend-lint
backend-lint:
	cd backend && golangci-lint run

# Format backend source code
.PHONY: backend-format
backend-format:
	cd backend && go fmt ./...

# Run backend tests
.PHONY: backend-test
backend-test:
	cd backend && go test ./...

# Run backend
.PHONY: backend-run
backend-run:
	cd backend/cmd/server && go run main.go

# Build backend
.PHONY: backend-build
backend-build:
	cd backend && go build -o bin/classroom cmd/server/main.go

# Run database
.PHONY: db-run
db-run:
	docker-compose up

# convert the backend link to an ngrok link
.PHONY: backend-ngrok
backend-ngrok:
	@echo ${PUBLIC_API_DOMAIN}
	cd backend && ngrok http --domain=${PUBLIC_API_DOMAIN} 8080
