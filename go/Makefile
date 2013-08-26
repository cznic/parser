.PHONY: all clean

all: editor
	go vet
	go install
	make todo

editor: clean parser.go
	go fmt
	go test -i
	go test
	go build

parser.go: parser.y
	go tool yacc -l -o $@ $<
	sed -i 's|//line .*:.*||' $@

todo:
	@grep -nr ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* *.go || true
	@grep -nr TODO *.go || true
	@grep -nr BUG *.go || true
	@grep -nr println *.go || true

clean:
	@go clean
	rm -f y.go

nuke: clean
	rm -f y.output
