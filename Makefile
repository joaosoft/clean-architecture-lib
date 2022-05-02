run:
	go run ./main.go

test:
	go test ./... -v

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*