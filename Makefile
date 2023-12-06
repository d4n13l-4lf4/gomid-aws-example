COV_FILE=$(PWD)/build/coverage.out

clean: 
	rm -rf $(PWD)/build

coverage:
	sh $(PWD)/scripts/go-coverage-test.sh --settings-file $(PWD)/scripts/go-coverage-test-settings.sh --check-coverage false
	go tool cover -html=$(COV_FILE)

lint: 
	find . -name \*.go ! -path "./vendor/*" -exec gofmt -w -l {} \;
	find . -name \*.go ! -path "./vendor/*" -exec goimports -w {} \;

build: 
	sh $(PWD)/scripts/go-build.sh --cmd-dir cmd --out-dir build

clear-mock:
	rm -rf mocks

gen-mock:
	mockery