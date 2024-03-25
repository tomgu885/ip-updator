GIT_COMMIT=`git rev-list --abbrev-commit --max-count=1 --all`
BUILD_TIME=`date +%FT%T%z`

help: ## show help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test:
	#ssh dev "pkill main-test"
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME} -s -w" -o bin/main-test main.go
	scp bin/main-test dev:/root/test/

test2:
	#ssh dev "pkill main-test"
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME} -s -w" -o bin/main-test main.go
	scp bin/main-test bwg:/root/test/


linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME} -s -w" -o bin/updator-amd64-`date +'%Y%m%d%H%M%S'` main.go