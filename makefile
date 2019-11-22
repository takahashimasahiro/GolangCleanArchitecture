build:
	docker-compose build
run:
	docker-compose run --rm app
up:
	docker-compose up
mysql:
	docker run -ti mysql bash
db-option:
	docker-compose run --rm db mysql --help

db-login:
	docker-compose run --rm db mysql -u root -p