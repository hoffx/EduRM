all: git_commit=$(shell git rev-list -1 HEAD)
all: version = development

all: deploy

deploy: deployMacOS

deployMacOS:
ifeq ($(version),)
	@echo "Assuming this is a development build ! You can specify a version using the following syntax: make deploy version=<version_string>"
endif
	@rm -r ./darwin/Contents/Resources/qml
	@cp -r ./qml ./darwin/Contents/Resources/qml
	@qtdeploy -ldflags "-X main.GitCommit=$(git_commit) -X main.Version=$(version) -X main.GuiPath=../Resources/qml/main.qml"