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
	go tool yacc -o $@ $<

todo:
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* *.go *.l || true
	@grep -n TODO *.go *.l || true
	@grep -n BUG *.go *.l || true
	@grep -n println *.go *.l || true

clean:
	@go clean
	rm -f y.output
