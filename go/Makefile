.PHONY: all clean

all: editor
	go vet
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
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* *.go *.y || true
	@grep -n TODO *.go *.y || true
	@grep -n BUG *.go *.y || true
	@grep -n println *.go *.y || true

clean:
	@go clean
	rm -f y.output
