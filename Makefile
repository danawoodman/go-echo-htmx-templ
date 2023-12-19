.PHONY: dev-server dev-tailwind dev-templ dev build-server build-tailwind build-templ build launch deploy

#-----------------------------------------------------
# DEV
#-----------------------------------------------------

dev:
	@make -j dev-server dev-templ dev-tailwind

dev-server:
	@gochange -k -i '**/*.go' '**/*.html' -- go run ./cmd/server

dev-templ:
	@templ generate --watch

dev-tailwind:
	@npm run dev

#-----------------------------------------------------
# BUILD
#-----------------------------------------------------

build:
	@make build-tailwind build-server build-templ

build-server:
	@go build -o bin/server ./cmd/server/main.go

build-templ:
	@templ generate

build-tailwind:
	@npm run build

#-----------------------------------------------------
# DATABASE
#-----------------------------------------------------

# drop:
# 	@go run ./cmd/drop

# seed:
# 	@go run ./cmd/seed

#-----------------------------------------------------
# DEPLOY
#-----------------------------------------------------

launch:
	@fly launch

deploy:
	@fly deploy

.DEFAULT_GOAL := dev  