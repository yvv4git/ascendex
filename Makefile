gen_mock:
	mockgen -source=connection.go -destination=mock/connection_mock.go -package mock

test:
	go test -v -p 1 ./...

build:
	go build -o ascendex example/main.go

run:
	./ascendex --symbol BTC_USDT --timeout 10s
