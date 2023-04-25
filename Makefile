PAKAGE_LIST := $(shell go list ./..)
shortURLz:
			go build -o shortURLz $(PAKAGE_LIST)
test:
			go test  -covermode=count -coverprofile=coverrage.out $(PAKAGE_LIST)
clean:
			rm -f shortURLz