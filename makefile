test_func=TestTemplateXSSServer
test:
	@go test -v -run=$(test_func)
