.PHONY: do-migrate
do-migrate:
	migrate -database 'postgres://root:123@localhost:5432/yandex?sslmode=disable' -path migrations up
