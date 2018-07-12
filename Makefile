SHELL:=/bin/bash
VERSION?=$(shell git describe --tags --always)
CURRENT_DOCKER_IMAGE=quintilesims/eks-sso:$(VERSION)
LATEST_DOCKER_IMAGE=quintilesims/eks-sso:latest

deps:
	go get github.com/golang/mock/mockgen/model
	go install github.com/golang/mock/mockgen

mocks:
	mockgen -package mock github.com/aws/aws-sdk-go/service/sts/stsiface STSAPI > pkg/mock/mock_sts.go && \
	sed -i '.tmp' s/github.com\\/quintilesims\\/eks-sso\\/vendor\\/// pkg/mock/mock_sts.go && \
	rm -f pkg/mock/mock_sts.go.tmp
	mockgen -package mock github.com/aws/aws-sdk-go/service/eks/eksiface EKSAPI > pkg/mock/mock_eks.go && \
	sed -i '.tmp' s/github.com\\/quintilesims\\/eks-sso\\/vendor\\/// pkg/mock/mock_eks.go && \
	rm -f pkg/mock/mock_eks.go.tmp
	mockgen -package mock github.com/quintilesims/eks-sso/pkg/auth Authenticator > pkg/mock/mock_authenticator.go
	mockgen -package mock github.com/quintilesims/eks-sso/pkg/auth AWS > pkg/mock/mock_aws.go
	mockgen -package mock github.com/quintilesims/eks-sso/pkg/auth Kubernetes > pkg/mock/mock_kubernetes.go

build:
	cd src/ui; npm run build;
	docker build -t $(CURRENT_DOCKER_IMAGE) .

release: build
	docker push $(CURRENT_DOCKER_IMAGE)
	docker tag $(CURRENT_DOCKER_IMAGE) $(LATEST_DOCKER_IMAGE)
	docker push $(LATEST_DOCKER_IMAGE)