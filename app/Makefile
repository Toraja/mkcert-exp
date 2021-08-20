SHELL = /bin/bash

.DEFAULT_GOAL := dummy

.PHONY: dummy
dummy:
	@echo Select a target.

.PHONY: run
run:
	@echo '--- Press Ctrl+C to stop the server ---'
	docker-compose exec app go run main.go
