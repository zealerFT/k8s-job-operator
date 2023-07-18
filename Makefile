.PHONY: ci-test
ci-test:
	@go test ./... -coverprofile .cover.txt
	@go tool cover -func .cover.txt
	@rm .cover.txt

.PHONY: build
build:
	rm -rf bin
	mkdir -p bin
	cd frontend && npm install && npm run build
	cd ..
	CGO_ENABLED=0 go build -o bin/k8s-job-operator -ldflags "-X k8s-job-operator/pod.appRelease=${release}" main.go
