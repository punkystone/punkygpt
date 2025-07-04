default: docker-build

docker-build:
	docker compose build

docker-up:
	docker compose up

docker-down:
	docker compose down