test_func=TestCookie
test:
	@go test -v -run=$(test_func)
