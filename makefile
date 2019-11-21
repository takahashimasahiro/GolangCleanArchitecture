build:
	docker-compose build
run:
	docker-compose run --rm app
up:
	docker-compose up
mysql:
	docker run -ti mysql bash