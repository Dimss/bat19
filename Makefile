objects:
	controller-gen object paths=./...

manifests:
	controller-gen crd paths=./... output:artifacts:config=./config/crd/bases

build-bat:
	go build -o bin/bat cmd/bat/*.go