SHELL:=/bin/bash
COV_FILE=$(PWD)/build/coverage.out

clean: 
	rm -rf $(PWD)/build

coverage:
	bash $(PWD)/scripts/go-coverage-test.sh --settings-file $(PWD)/scripts/go-coverage-test-settings.sh --check-coverage false
	go tool cover -html=$(COV_FILE)

lint: 
	find . -name \*.go ! -path "./vendor/*" -exec gofmt -w -l {} \;
	find . -name \*.go ! -path "./vendor/*" -exec goimports -w {} \;
	golangci-lint run
	terraform fmt .

docs:
	terraform-docs markdown table . > terraform.md

build:
	bash $(PWD)/scripts/go-build.sh --cmd-dir cmd --out-dir build

clear-mock:
	rm -rf mocks

gen-mock:
	mockery