.DEFAULT_GOAL: up

.PHONY:
up:
	make docker/up
	sleep 2
	make migrate/up

include make/docker.mk

include make/migrate.mk

include make/golang.mk