GOPATH=/Users/avmalov/go

build:
	GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose
	GO111MODULE=off GOOS=linux GOARCH=amd64 go build -i -o bin/goose ${GOPATH}/src/github.com/pressly/goose/cmd/goose
	GOOS=linux GOARCH=amd64 go build -i -o bin/bot cmd/main.go

migrate:
	cd migrations && \
	goose postgres "user=tiger dbname=tiger password=tiger host=127.0.0.1 sslmode=disable" up