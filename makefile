test_func=TestDownloadFile
test:
	@go test -v -run=$(test_func)
