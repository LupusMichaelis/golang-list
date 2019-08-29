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

run: $(SRCS)
	go $@ .

test: $(TESTS) $(SRCS)
	go test -v ./...

format: $(SRCS) $(TESTS)
	go fmt ./...

host:
	-@ln -s $(PWD) $(HOME)/.local/go-lang-root/src/$(VENDOR)/
