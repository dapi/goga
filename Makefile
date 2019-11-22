.PHONY: test
test:
	richgo test -v ./...

goga:
	go build ./goga.go
