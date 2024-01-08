lint-version:
	golangci-lint --version
lint-run:
	golangci-lint run
build:
	go build main
run:
	./main -source /home/useralex/bashwrapper/bashwrapper/folder/source -destination /home//useralex/bashwrapper/bashwrapper/folder/destination
