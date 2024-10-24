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

# Run backend
.PHONY: db-run
db-run:
	docker-compose up
