postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=etu28 -d postgres:14-alpine

create_db:
	soda create

drop_db:
	docker exec -it postgres14 dropdb tasks-management-with-go_development

migrate_up:
	soda migrate up -p ./internal/persistence/migrations

migrate_down:
	soda migrate down -p ./internal/persistence/migrations

.PHONY: postgres create_db drop_db migrate_up migrate_down
