run:
	@go run main.go api

coverage: ## generate and display code coverage report
	@echo "Generating coverage report..."
	@go test -v -coverprofile=coverage.out ./...
	@echo ""
	@echo "Coverage summary:"
	@go tool cover -func=coverage.out | tail -1
	@echo ""
	@echo "Generating HTML coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated: coverage.html"

coverage-view: ## generate and open code coverage report in browser
	@echo "Generating coverage report..."
	@go test -v -coverprofile=coverage.out ./...
	@echo ""
	@echo "Coverage summary:"
	@go tool cover -func=coverage.out | tail -1
	@echo ""
	@echo "Opening coverage report in browser..."
	@go tool cover -html=coverage.out -o coverage.html
	@open coverage.html
	
.PHONY: coverage coverage-view test-coverage-min