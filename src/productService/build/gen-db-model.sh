#!/bin/bash

# To be able to use the protobufs as domain for Gorm ORM
domaindir=internal/domain
protoc-go-inject-tag -input=$domaindir/store.pb.go
