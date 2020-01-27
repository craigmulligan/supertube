.PHONY: install 
install: go.sum
	go get -v -t

.PHONY: build 
build: install
	go build -o supertube 

.PHONY: test
test: build
	go test	

.PHONY: dist
dist:
	./dist-build.sh github.com/hobochild/supertube
