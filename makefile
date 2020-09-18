all: go.mod
	go get -u

go.mod:
	go mod init