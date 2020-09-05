# use bash as the default shell to be used in the recipe
SHELL := /bin/bash

gen-proto :
		# Generate the stubs from protobuf to Go
		@echo generate the Go code from protobuf
		$(SHELL) build/protogen.sh

gen-models :
		# Convert the generate Go stubs to database models
		@echo creating database models from the *.pb.go file
		$(SHELL) build/gen-db-model.sh

		# Generate the mocks for productService
gen-mocks :
		@echo generating mocks for tests
		$(SHELL) build/mock-gen.sh
