.PHONY: test fmt lint

test:
	@echo "🧪 Running tests..."
	go test ./...

fmt:
	@echo "📝 Formatting code..."
	gofmt -w -l .

lint:
	@echo "🔍 Linting code..."
	go vet ./...
	go tool staticcheck ./...
