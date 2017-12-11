#
# Makefile
# @author Evgeny Ukhanov <mrlsd@ya.ru>
#


run-server:
	@go build server.go && ./server

build-server:
	@go build server.go

run-client:
	@go build client.go && ./client

build-client:
	@go build client.go


fmt:
	@gofmt -w -l .

test:
	@go test -v ./...
