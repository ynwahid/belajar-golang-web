test_func=Foo
test:
	@go test -v -run=$(test_func)
