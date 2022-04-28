run:
	go run ./main.go

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*