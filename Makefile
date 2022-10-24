dev_start:
	docker-compose -f ./docker/docker-compose.yml  up -d --build --remove-orphans

dev_stop:
	docker-compose -f ./docker/docker-compose.yml  down

# golang-migrate must be installed (brew install golang-migrate)
db_setup:
	docker exec timescale_db psql -U postgres -d homework -c "\COPY cpu_usage FROM /docker-entrypoint-initdb.d/cpu_usage.csv CSV HEADER"


test:
	cd src;\
	go test -v -cover ./...

build:
	cd src;\
	go build -o ../docker/bin/timescale-take-home-assignment

run:
	cd src;\
	go run main.go