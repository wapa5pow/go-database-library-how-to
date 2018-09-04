up:
	docker-compose up

foreman:
	foreman start

goose_postgres:
	cd goose && goose postgres "user=postgres dbname=postgres sslmode=disable password=password" $(arg)

goose_mysql:
	cd goose && goose mysql "username:password@/sample?parseTime=true" $(arg)

