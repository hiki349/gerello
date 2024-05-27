DATE := $(shell date +%s%3N)

define create_migration
	$(eval FILE_NAME := $(DATE)_$(name))
	$(eval EXT := ".sql")

	@echo "" > migrations/$(FILE_NAME)_up$(EXT)
	@echo "created $(FILE_NAME)_up$(EXT)"
	@echo "" > migrations/$(FILE_NAME)_down$(EXT)
	@echo "created $(FILE_NAME)_down$(EXT)"
endef

run:
	@go run cmd/main/main.go

create-migration:
	$(create_migration)

docker:
ifeq ($(cmd),up)
	@docker compose up -d
else
	@docker compose down
endif
