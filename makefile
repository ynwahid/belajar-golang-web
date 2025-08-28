test_func=TestTemplateAutoEscapeServer
test:
	@go test -v -run=$(test_func)
