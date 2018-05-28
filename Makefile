
all: udp_receive_linux udp_receive_osx



udp_receive_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o udp_receive_linux

udp_receive_osx:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -o udp_receive_osx

run:
	go run server.go --host=router.eu.thethings.network:1700

test:
	go test -v

.PHONY: all run test
