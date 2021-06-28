SHELL:=/bin/bash

.PHONY:build run clean

build:
	go build -o doraemon cmd/alter-gateway/main.go

run:
	go build -o doraemon cmd/alter-gateway/main.go
	./doraemon -c configs/alterGateway.yml

clean:
	rm logs doraemon -rf
