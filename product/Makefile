# use bash as the default shell to be used in the recipe
SHELL := /bin/bash
.PHONY : gen-proto # to make make not be confused by the build directory

gen-proto :
		@echo generate the Go code from protobuf
		$(SHELL) build/proto-gen.sh

gen-mocks :
		@echo generating mocks for tests
		$(SHELL) build/mock-gen.sh
