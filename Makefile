dev-start:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build
dev-stop:
	docker-compose down

prod-start:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build --detach
prod-stop:
	docker-compose down

debug:
	~/go/bin/dlv 
rm-volumes:
	docker-compose down --volumes
