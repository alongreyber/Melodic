dev:
	docker-compose up --build
prod:
	docker-compose up -f docker-compose.yaml -f docker-compose.prod.yaml --build 
debug:
	~/go/bin/dlv 

prisma-deploy:
	export PRISMA_MANAGEMENT_API_SECRET=my-secret-0000 
	prisma deploy
	prisma generate
