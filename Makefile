#!make
include .env

DB_PATH=pkg/database/migrations

install:
	@echo 'Install dependencies'
	go mod tidy

build:
	@echo 'Build app'
	go build cmd/main.go

run:
	@echo 'Run app'
	go run cmd/main.go

migrateup:
ifneq (undefined, $(filter $(DB_USER),$(DB_PASS),$(DB_NAME),$(DB_HOST),$(DB_PORT)))
	@echo 'Run database migrations'
	migrate -path $(DB_PATH) -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up
else
	@echo 'Please define in .env file variables DB_USER, DB_PASS, DB_NAME, DB_HOST, DB_PORT'
endif

migratedown:
ifneq (undefined, $(filter $(DB_USER),$(DB_PASS),$(DB_NAME),$(DB_HOST),$(DB_PORT)))
	@echo 'Revert database migration file'
	migrate -path $(DB_PATH) -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down
else
	@echo 'Please define in .env file variables DB_USER, DB_PASS, DB_NAME, DB_HOST, DB_PORT'
endif
	