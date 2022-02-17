objects:
	controller-gen object paths=./...
manifests:
	controller-gen crd paths=./... output:artifacts:config=./config/crd/bases
build-t:
	go build -o bin/t cmd/template/*.go