
# cat -e -t -v Makefile

.DEFAULT_GOAL := buildAndRun

build:
	@echo "go build..."
	qtdeploy build desktop

run:
	./deploy/linux/FenixTestInstructionBuilder

buildAndRun:
	qtdeploy build desktop
	./deploy/linux/FenixTestInstructionBuilder

installStep0:
	go mod init

installStep1:
	go mod download && go get -u -v github.com/therecipe/qt/cmd/qtsetup && go get -u -v github.com/therecipe/qt/cmd/...

installStep2:
	go mod vendor && cd ./vendor/github.com/therecipe && git clone https://github.com/therecipe/env_$(go env GOOS)_amd64_512.git && cd ../../..