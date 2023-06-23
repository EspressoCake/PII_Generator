BIN := PII_Generator

.PHONY: preamble
.SILENT: preamble
preamble:
	@rm go.mod 2>/dev/null || true
	@rm go.sum 2>/dev/null || true

.PHONY: clean
clean: preamble
	@rm bins/* 2>/dev/null || true

.PHONY: all
all: clean preamble
	@go mod init ${BIN}
	@go mod tidy
	@GOOS=linux go build -o bins/${BIN}_linux main.go
	@GOOS=darwin go build -o bins/${BIN}_darwin main.go
	@GOOS=windows go build -o bins/${BIN}_windows.exe main.go