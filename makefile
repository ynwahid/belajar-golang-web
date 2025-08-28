test_func=TestRedirect
test:
	@go test -v -run=$(test_func)
