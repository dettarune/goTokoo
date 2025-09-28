APP_NAME=goTokoo
BINARY=bin/$(APP_NAME)

# Build binary
build:
	go build -o $(BINARY) ./cmd/server/main.go

# Run server
run:
	go run ./cmd/server/main.go

# Jalankan binary yang udah dibuild
start: build
	./$(BINARY)

# Clean binary
clean:
	rm -rf $(BINARY)

# Jalankan unit test
test:
	go test ./... -v

# Contoh migrate database (kalau nanti dipake)
migrate:
	go run ./cmd/migrate.go

