.PHONY: dbup create up dw migrate psql

dbup:
	docker run --name profitti-db -p 5434:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:alpine

create:
	docker exec -it profitti-db createdb --username=postgres --owner=postgres profitti

up:
	goose up

dw:
	goose down

migrate:
	goose create $(name) sql

psql:
	docker exec -it profitti-db psql -U postgres -d profitti
