.PHONY: generate
generate:
	go generate ./pkg/...

.PHONY: build
build:
	go build -tags dev -o vfserver github.com/gargath/vfsgen-example/cmd

.PHONY: release
release: generate
	go build -o vfserver github.com/gargath/vfsgen-example/cmd

.PHONY: clean
clean:
	find . -name \*vfsdata.go -exec rm -f {} \;
	rm -f vfserver
