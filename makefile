test_func=TestFileServer
test:
	@go test -v -run=$(test_func)
