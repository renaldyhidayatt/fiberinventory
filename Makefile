run:
	go run cmd/app/main.go


gen:
	mockgen -source=internal/service/interfaces.go -destination=internal/service/mocks/mock.go
	mockgen -source=internal/repository/interfaces.go -destination=internal/repository/mocks/mock.go

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out | grep "total"