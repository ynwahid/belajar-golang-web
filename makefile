test_func=TestServeFileServer
test:
	@go test -v -run=$(test_func)
