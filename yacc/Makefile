.PHONY: all clean

all: editor
	go vet || true
	go install
	make todo

editor: parser.go
	go fmt
	go test -i
	go test
	go build

parser.go: parser.y
	go tool yacc -l -o $@ $<
	sed -i 's|//line .*:.*||' $@

todo:
	@grep -nr ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* *.go *.y || true
	@grep -nr TODO *.go *.y || true
	@grep -nr BUG *.go *.y || true
	@grep -nr println *.go *.y || true

clean:
	@go clean
	rm -f y.output
