
stop:
	docker-compose down

run:
	docker-compose up -d


migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:Merei04977773@@localhost:5434/kazakTeam?sslmode=disable' up

start:
	go run ./cmd/main.go


