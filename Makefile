# Define your services
SERVICES=db backend
IMAGE=khoury-classroom-plugin-backend:latest

# Build and run the backend
.PHONY: backend
backend:
	@echo "Starting backend containers..."
	docker compose up $(DETACHED)

# Build a fresh instance of the backend
.PHONY: restart-backend
restart-backend:
	@echo "Stopping and removing containers..."
	docker rm -f $(SERVICES)
	@echo "Removing backend image..."
	docker rmi $(IMAGE)
	@echo "Starting fresh instance of backend containers..."
	docker compose up --build $(DETACHED)

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
all:
	$(MAKE) DETACHED=-d backend
	$(MAKE) frontend