test_func=TestFileServerGolangEmbed
test:
	@go test -v -run=$(test_func)
