.PHONY: do-migrate
do-migrate:
	migrate -database 'postgres://root:123@localhost:5432/yandex?sslmode=disable' -path migrations up

#migrate create -ext sql -dir migrations -seq init_teams_table