PAKAGE_LIST := $(shell go list ./..)
shortURLz:
			go build -o shortURLz $(PAKAGE_LIST)
test:
			go test $(PAKAGE_LIST)
clean:
			rm -f shortURLz