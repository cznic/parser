# Copyright 2013 The parser Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY: all clean

all: editor
	go vet || true
	golint .
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
