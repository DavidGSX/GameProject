#!/bin/sh
protoc --go_out=. *.proto

sed "s/package dbProto/package main/" genDB.go > genDBTemp.go
go build genDBTemp.go
chmod a+x genDBTemp
./genDBTemp
rm genDBTemp.go
rm genDBTemp
