
downdb:
	docker-compose down

updb:
	docker-compose up -d


migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:Merei04977773@@localhost:5434/shop?sslmode=disable' up

start:
	go run ./cmd/main.go


