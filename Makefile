# Define your services
SERVICES=db backend
IMAGE=khoury-classroom-plugin-backend:latest

# Build and run backend
.PHONY: backend
backend:
	@echo "Stopping and removing containers..."
	docker rm -f $(SERVICES)
	@echo "Removing backend image..."
	docker rmi $(IMAGE)
	@echo "Starting fresh instance of backend containers..."
	docker compose up --build -d

# Build and run frontend
.PHONY: frontend
frontend:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install
	@echo "Linting frontend source code..."
	cd frontend && npx eslint . || { echo "Linting failed. Exiting."; exit 1; }
	@echo "Linting passed. Starting frontend..."
	cd frontend && npm run dev

# Build and run whole app
.PHONY: all
all: backend frontend