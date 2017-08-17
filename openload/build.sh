#!/bin/bash
GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o release/openload_linux_amd64 ../main.go
echo "编译openload_linux_amd64完成"
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/openload_windows_amd64 ../main.go
echo "编译openload_windows_amd64完成"
GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o release/openload_mac_amd64 ../main.go
echo "编译openload_mac_amd64完成"
GOOS=linux   GOARCH=386 CGO_ENABLED=0 go build -o release/openload_linux_386 ../main.go
echo "编译openload_linux_386完成"
GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o release/openload_windows_386 ../main.go
echo "编译openload_windows_386完成"
GOOS=darwin  GOARCH=386 CGO_ENABLED=0 go build -o release/openload_mac_386 ../main.go
echo "编译openload_mac_386完成"
