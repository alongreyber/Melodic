dev:
	docker-compose up --build
prod:
	docker-compose up -f docker-compose.yaml -f docker-compose.prod.yaml --build 
debug:
	~/go/bin/dlv 
rm-volumes:
	docker-compose down --volumes
