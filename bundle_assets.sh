#!/bin/sh
EECSPROJECT_PATH=$GOPATH/github.com/moosingin3space/eecsproject
cd $EECSPROJECT_PATH
mkdir internal
go-bindata -o temp.go assets/...
sed 's/package main/package internal/' <temp.go >internal/data.go
rm temp.go
