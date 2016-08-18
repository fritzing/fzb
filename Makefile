test: test-src test-bin

test-src:
	@echo "TEST src"
	@echo "---------------------------------------------------------"
	@cd src/go && go test -v ./...

test-bin:
	@echo "TEST bin"
	@echo "---------------------------------------------------------"
	@echo "build the commandline tool"
	@cd bin/fzb && go build
	@echo "run the cli..."
	./bin/fzb/fzb src/fixture/test16.fzb
