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
	@$(eval CMD=$(filter-out $@,$(MAKECMDGOALS)))
	@go run $(MAIN_GO) init $(CMD)

# Exec target with arguments
.PHONY: exec
exec:
	@$(eval CMD=$(filter-out $@,$(MAKECMDGOALS)))
	@go run $(MAIN_GO) exec $(CMD)

# Build target
.PHONY: build
build:
	@go build -o $(OUTPUT_BINARY)

# make move DEST=~/sandbox/express-dev-docker
move:
	@make build && mv -f chalet $(DEST)
