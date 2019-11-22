.PHONY: test
test:
	richgo test -v ./...

goga:
	go build ./goga.go

test_add:
	go run ./goga.go add https://github.com/dapi/elements/blob/master/spinner.js
