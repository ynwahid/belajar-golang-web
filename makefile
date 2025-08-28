test_func=TestMiddleware
test:
	@go test -v -run=$(test_func)
