pull:
	docker pull mariadb:10.7.4

db:
	docker-compose -f ./docker/docker-compose.yml up -d --build db

start-db:
	docker-compose -f ./docker/docker-compose.yml start db

stop-db:
	docker-compose -f ./docker/docker-compose.yml stop db

rm-db: stop-db
	docker-compose -f ./docker/docker-compose.yml rm db
	docker rmi -f $(docker images -f "dangling=true" -q)
