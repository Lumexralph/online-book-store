#!/bin/bash

# To be able to use the protobufs as domain for Gorm ORM, run the following commands
# in the terminal in the directory where the stub was generated
domaindir=internal/domain
protoc-go-inject-tag -input=$domaindir/store.pb.go
ls