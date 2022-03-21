
test:
	go test

testv:
	go test -v

COVERFILE="coverage.txt"

cover:
	go test -cover -coverprofile $(COVERFILE)
	go tool cover -html=$(COVERFILE)


.PHONY: test testv cover


