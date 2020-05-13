# use bash as the default shell to be used in the recipe
SHELL := /bin/bash
.PHONY : build # to make make not be confused by the build directory

build :
		@echo generate the Go code from protobuf
		$(SHELL) build/proto-gen.sh

