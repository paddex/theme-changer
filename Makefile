COMMIT := $(shell git log -1 --format="%h")
DATE := $(shell date "+%F %T")

build:
	@go build -o bin/theme -ldflags="-X 'paddex.net/theme-changer/cmd.commit=${COMMIT}' -X 'paddex.net/theme-changer/cmd.date=${DATE}'"

install:
	@go install

clean:
	@rm -r bin
