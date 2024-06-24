# Variables
MAIN_GO := main.go
OUTPUT_BINARY := chalet

# Default target
.PHONY: null
null:
	@go run $(MAIN_GO)

# Run target
.PHONY: run
run:
	@go run $(MAIN_GO) run

# Init target
.PHONY: init
init:
	@go run $(MAIN_GO) init

# Exec target with arguments
.PHONY: exec
exec:
	@$(eval CMD=$(filter-out $@,$(MAKECMDGOALS)))
	@go run $(MAIN_GO) exec $(CMD)

# Build target
.PHONY: build
build:
	@go build -o $(OUTPUT_BINARY)