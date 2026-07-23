.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build:
	sudo docker build -t Suiren91/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local:
	sudo docker compose build --no-cache

up:
	sudo docker compose up -d

down:
	sudo docker compose down

logs:
	sudo docker compose logs -f

ps:
	sudo docker compose ps

test:
	go test -race -shuffle=on ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n",$$1, $$2}'
