# use bash as the default shell to be used in the recipe
# TODO: Write script to create the client collection in mongodb
SHELL := /bin/bash
.PHONY : build # to make make not be confused by the build directory

build :
		@echo generate the Go code for client proto
		$(SHELL) build/proto-gen.sh

