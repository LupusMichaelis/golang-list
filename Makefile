.PHONY: test format

SRCS=\
	main.go \
	list.go \

run: $(SRCS)
	go $@ $?

format: $(SRCS)
	go fmt $?
