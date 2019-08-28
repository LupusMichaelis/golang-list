.PHONY: test format

SRCS=\
	main.go \

run: $(SRCS)
	go $@ $?

format: $(SRCS)
	go fmt $?
