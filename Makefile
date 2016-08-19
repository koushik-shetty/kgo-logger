#binary out put dir
OUT_DIR = out/

#directories containing the source files. 
SRC_FILES = logger.go

#Default make command.
all: get-deps build test

ver: 
	@go version

build: ver $(SRC_FILES)
	@echo "building library"
	@go build logger.go
	@echo "build command complete"

test: ver
	@echo "Starting go tests.."
	go test
	@echo "tests completed."

setup:
	@echo "preparing project..."
	./ci_scripts/pre_install.sh
	@echo "project prepared.\nRun 'make' to build the project"

get-deps:
	@echo "installing dependencies..."
	@$(GOPATH)/bin/glide install