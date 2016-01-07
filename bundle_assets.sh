#!/bin/sh
mkdir internal
go-bindata -o temp.go assets/...
sed 's/package main/package internal/' <temp.go >internal/data.go
rm temp.go
