build:
	docker-compose build
run:
	docker-compose run --rm app
up:
	docker-compose up
mysql:
# https://stackoverflow.com/questions/23234379/installing-mysql-in-docker-fails-with-error-message-cant-connect-to-local-mysq
	docker-compose run db bash
# db-option:
# 	docker-compose run --rm db mysql --help

db-login:
	# docker-compose run --rm db mysql -u root -p
	docker run -ti db -uroot