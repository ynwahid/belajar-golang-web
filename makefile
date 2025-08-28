test_func=TestTemplateAutoEscapeDisabled
test:
	@go test -v -run=$(test_func)
