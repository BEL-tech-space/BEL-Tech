run:
	@cd frontend && $(MAKE) -f Makefile build
	@go run .

PHONY: run