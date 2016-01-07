#!/bin/sh
go-bindata -o temp.go assets/...
sed 's/package main/package internal/' <temp.go >internal/data.go
rm temp.go
