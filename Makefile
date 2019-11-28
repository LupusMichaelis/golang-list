.PHONY: \
	format \
	host
	run \
	test \

VENDOR=lupusmic.org

SRCS=\
	main.go \
	list/list.go \

TESTS=\
	list/list_test.go \

run: host $(SRCS)
	go $@ .

test: host $(TESTS) $(SRCS)
	go test -v ./...

format: $(SRCS) $(TESTS)
	go fmt ./...

host:
	test -h $(GOPATH)/src/$(VENDOR)/$(notdir $(PWD)) \
		|| ln -s $(PWD) $(GOPATH)/src/$(VENDOR)/
	go get -t ./...
