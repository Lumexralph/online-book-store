# use bash as the default shell to be used in the recipe
SHELL := /bin/bash
.PHONY : gen-js-proto # to make make not be confused by the build directory

gen-js-proto :
		@echo generate the gRPC-web Javascript code from product service protobuf
		$(SHELL) build-js-web.sh