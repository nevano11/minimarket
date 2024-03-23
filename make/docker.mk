.PHONY: docker/down/v
docker/down/v:
	docker compose down -v

.PHONY: docker/down
docker/down:
	docker compose down

.PHONY: docker/up
docker/up:
	docker compose up -d --build