OUT_DIR := ./out

GO_FILES := $(shell find . -type f -name '*.go')
GO_MODULE := github.com/eikendev/gitea-issue-exporter

.PHONY: build
build:
	mkdir -p $(OUT_DIR)
	go build -ldflags "-w -s" -o $(OUT_DIR)

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)

.PHONY: test
test:
	if [ -n "$$(gofumpt -l $(GO_FILES))" ]; then echo "Code is not properly formatted"; exit 1; fi
	if [ -n "$$(goimports -l -local $(GO_MODULE) $(GO_FILES))" ]; then echo "Imports are not properly formatted"; exit 1; fi
	go vet ./...
	misspell -error $(GO_FILES)
	gocyclo -over 10 $(GO_FILES)
	staticcheck ./...
	errcheck -ignoregenerated ./...
	gocritic check -disable='#experimental,#opinionated' -@ifElseChain.minThreshold 3 ./...
	revive -set_exit_status -exclude $(EXCLUDE_GO_DIRS) ./...
	go test -v -cover ./...
	gosec -exclude-dir=tests -exclude-generated ./...
	govulncheck ./...
	@printf '\n%s\n' "> Test successful"

.PHONY: setup
setup:
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/mgechev/revive@latest
	go install github.com/ogen-go/ogen/cmd/ogen@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: format
format:
	gofumpt -l -w .
	goimports -w -local $(GO_MODULE) $(GO_FILES)
