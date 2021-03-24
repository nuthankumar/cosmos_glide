BUILDPATH=$(CURDIR)
GO=$(shell which go)
GLIDE=$(shell which glide)
GLIDEUP=$(GLIDE) up --skip-test
GLIDEINSTALL=$(GLIDE) install
GOINSTALL=$(GO) install
GOBUILD=$(GO) build  -o ./bin/ccspaas_crd ./pkg/controller/*.go
GORUN=$(GO) run *.go
GOCLEAN=$(GO) clean
GOGET=$(GO) get -v
DOCKER=$(shell which docker)
DOCKERBUILD=$(DOCKER) build --build-arg HTTP_PROXY=http://3.20.109.241:88 --build-arg HTTPS_PROXY=http://3.20.109.241:88 -t hc-eu-west-aws-artifactory.cloud.health.ge.com/docker-snapshot-clinical-care-app/ccs-paas/scylla-crd:latest .
DOCKERPUSH=$(DOCKER) push hc-eu-west-aws-artifactory.cloud.health.ge.com/docker-snapshot-clinical-care-app/ccs-paas/scylla-crd:latest

export GOPATH=$(CURDIR)
export GOBIN=$(CURDIR)/bin
export KUBECONFIG=/root/.kube/config

shell:
	@$(SHELL)
makedir:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/src ] ; then mkdir -p $(BUILDPATH)/src ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi
	@echo "building tree completed"
get:
	@$(GLIDEUP)
install:
	@$(GLIDEINSTALL)
run:
	@echo "running controller"
	@$(GORUN)

build:
	@echo "start go building..."
	$(GOBUILD)
	@echo "build completed ! DONE!"
docker:
	@echo "building dockerfile"
	$(DOCKERBUILD)
	@echo "docker image built"
	@echo "pushing docker image"
	$(DOCKERPUSH)
	@echo "docker image pushed"
