
proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc -I/usr/local/include -I. \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:. proto/*.proto

test:
	@rm -rf mock
	@mkdir mock
	mockgen -source=service/userservice.go -package=mock > mock/mock_userrepo.go
	go test service/userservice_test.go

.PHONY: proto test