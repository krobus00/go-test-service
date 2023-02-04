launch_args=
test_args=-coverprofile cover.out && go tool cover -func cover.out
cover_args=-cover -coverprofile=cover.out `go list ./...` && go tool cover -html=cover.out

.PHONY: build

tidy:
	go mod tidy

lint:
	golangci-lint run

build:
	# build binary file
	go build -ldflags "-s -w" -o go-test-service main.go
ifeq (, $(shell which upx))
	$(warning "upx not installed")
else
	# compress binary file if upx command exist
	upx -9 go-test-service
endif

run:
ifeq (server, $(filter server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
else ifeq (worker, $(filter worker,$(MAKECMDGOALS)))
	$(eval launch_args=worker $(launch_args))
endif
	./go-test-service $(launch_args)

test:
ifeq (, $(shell which richgo))
	go test ./... $(test_args)
else
	richgo test ./... $(test_args)
endif
	

cover: test
ifeq (, $(shell which richgo))
	go test $(cover_args)
else
	richgo test $(cover_args)
endif

%:
	@: