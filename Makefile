dev:
	docker-compose up --build
prod:
	docker-compose up -f docker-compose.yaml -f docker-compose.prod.yaml --build 
debug:
	~/go/bin/dlv 

VOLUMES = $(shell docker volume ls -q)
remove-volumes:
	docker-compose down --volumes
	docker volume rm $(VOLUMES)
