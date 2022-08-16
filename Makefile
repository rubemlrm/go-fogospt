.PHONY: help
# Display all commands
help: # Show this help.
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

.PHONY: generate-mocks
# Generate mocks
generate-mocks:
	grep -rl  internal/ -e 'type \s*[a-zA-Z0-9]\+\s* interface' | xargs -n 1 dirname | uniq | xargs -P 4 -I % sh -c 'mockery -name "[a-zA-Z0-9]*" -note "Run \`make generate-mocks\` to regenerate this file." -outpkg mocks -dir "%" -output "%/mocks"'

.PHONY: tests
# Run golang tests
tests:
	go test -v ./...


.PHONY: tests-coverage
# Get test converages
tests-coverage:
	go test -coverprofile=coverage.out ./fogospt
	go tool cover -html=coverage.out
