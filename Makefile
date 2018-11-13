# Makefile for the repo.
# This is used to help coordinate some of the tasks around the build process.
# Most of the actual build steps should be managed by Bazel.

## Bazel command to use.
BAZEL     := bazel

## Dep command to use.
DEP       := dep

## Minikube command to use.
MINIKUBE  := minikube

## Kubectl command to use.
KUBECTL := kubectl

## The directory to write template files to for skaffold (with respect to bazel info workspace).
SKAFFOLD_DIR := $$(bazel info workspace)/skaffold_build

.PHONY: clean
clean:
	$(BAZEL) clean 
	rm -rf $(SKAFFOLD_DIR)

.PHONY: pristine
pristine:
	$(BAZEL) clean --expunge

.PHONY: build
build: ## Run the full build (except UI).
	$(BAZEL) build //...

.PHONY: test
test: ## Run all the tests (except UI).
	$(BAZEL) test //...

.PHONY: dep-ensure
dep-ensure: ## Ensure that go dependencies exist.
	$(DEP) ensure

gazelle-repos: Gopkg.lock
	$(BAZEL) run //:gazelle -- update-repos -from_file=Gopkg.lock

gazelle: gazelle-repos ## Run gazelle to update go build rules.
	$(BAZEL) run //:gazelle

go-setup: dep-ensure gazelle

k8s-load-certs:
	-$(KUBECTL) delete secret custom-tls-cert
	$(KUBECTL) create secret tls custom-tls-cert --key services/certs/server.key --cert services/certs/server.crt

k8s-load-dev-secrets: #Loads the secrets used by the dev environment. At some point it might makse sense to move this into a dev setup script somewhere.
	-$(KUBECTL) delete secret pl-app-secrets
	$(KUBECTL) create secret generic pl-app-secrets \
		--from-literal=jwt-signing-key=ABCDEFG \
		--from-literal=session-key=test-session-key \
		--from-literal=auth0-client-id=qaAfEHQT7mRt6W0gMd9mcQwNANz9kRup \
		--from-literal=auth0-client-secret=_rY9isTWtKgx2saBXNKZmzAf1y9pnKvlm-WdmSVZOFHb9OQtWHEX4Nrh3nWE5NNt

dev-env-start: ## Start dev environment.
	$(MINIKUBE) start --cpus 6 --memory 8192 --vm-driver hyperkit --mount-string="$(HOME):$(HOME)" --mount
	$(MAKE) k8s-load-certs
	$(MAKE) k8s-load-dev-secrets

dev-docker-start:
	@eval $$(minikube docker-env); ./scripts/run_docker.sh --extra_args="$(DEV_DOCKER_EXTRA_ARGS)"

dev-env-stop: ## Stop dev environment.
	$(MINIKUBE) stop

dev-env-teardown: dev-env-stop ## Clean up dev environment.
	$(MINIKUBE) delete

skaffold-dev: ## Run Skaffold in the dev environment.
	$(BAZEL) run //templates/skaffold:skaffoldtemplate -- --build_dir $(SKAFFOLD_DIR)

skaffold-prod: ## Run Skaffold in the prod environment.
	$(BAZEL) run //templates/skaffold:skaffoldtemplate -- --build_dir $(SKAFFOLD_DIR) --prod

skaffold-staging: ## Run Skaffold in the staging environment.
	$(BAZEL) run //templates/skaffold:skaffoldtemplate -- --build_dir $(SKAFFOLD_DIR) --staging

help: ## Print help for targets with comments.
	@echo "Usage:"
	@echo "  make [target...] [VAR=foo VAR2=bar...]"
	@echo ""
	@echo "Useful commands:"
# Grab the comment prefixed with "##" after every rule.
	@grep -Eh '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
		sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(cyan)%-30s$(term-reset) %s\n", $$1, $$2}'
	@echo ""
	@echo "Useful variables:"
# Grab the comment prefixed with "##" before every variable.
	@awk 'BEGIN { FS = ":=" } /^## /{x = substr($$0, 4); \
    getline; if (NF >= 2) printf "  $(cyan)%-30s$(term-reset) %s\n", $$1, x}' $(MAKEFILE_LIST) | sort
	@echo ""
	@echo "Typical usage:"
	@printf "  $(cyan)%s$(term-reset)\n    %s\n\n" \
		"make build" "Run a clean build and update all the GO deps." \
		"make pristine" "Delete all cached builds." \
