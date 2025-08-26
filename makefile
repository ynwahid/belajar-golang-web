test_func=TestServeFileEmbed
test:
	@go test -v -run=$(test_func)
