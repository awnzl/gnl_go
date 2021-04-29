GOTEST	:= go test

GOBASE	:= $(shell pwd)
TESTS	:= $(GOBASE)/pkg/*

.PHONY: test
test:
	$(GOTEST) $(TESTS)
