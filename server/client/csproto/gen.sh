#!/bin/sh
protoc --go_out=. *.proto

sed "s/package csproto/package main/" genMsg.go > genMsgTemp.go
go build genMsgTemp.go
chmod a+x genMsgTemp
./genMsgTemp
rm genMsgTemp.go
rm genMsgTemp
