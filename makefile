test_func=TestUploadForm
test:
	@go test -v -run=$(test_func)
