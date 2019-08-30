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
	go get -t ./...
	test -h $(GOROOT)/src/$(VENDOR)/$(notdir $(PWD)) \
		|| ln -s $(PWD) $(GOROOT)/src/$(VENDOR)/
