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
	./bin/fzb/fzb -version
	# ./bin/fzb/fzb -f src/fixture/test16.fzb

install:
	cd bin/fzb && go install

cover:
	cd src/go && go test -coverprofile cover.out
	cd src/go && go tool cover -html=cover.out -o cover.html

