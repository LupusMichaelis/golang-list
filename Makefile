.PHONY: test format

SRCS=\
	main.go \
	list.go \

TESTS=\
	list_test.go \

run: $(SRCS)
	go $@ $?

test: $(TESTS) $(SRCS)
	go test -v

format: $(SRCS) $(TESTS)
	go fmt $?
