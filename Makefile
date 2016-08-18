test: test-src test-bin

test-src:
	@echo "test the src"
	@cd src/go && go test -v

test-bin:
	@echo "build the commandline tool"
	@cd bin/fzb && go build
	@echo "run the cli..."
	./bin/fzb/fzb src/fixture/test1.fzb
