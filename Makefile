run:
	go run .

dev:
	air

dev-deps:
	go install github.com/cosmtrek/air@latest

deps:
	go get -u github.com/gin-gonic/gin

build:
	go build -o tmp/ginie .

clean:
	rm -fr tmp


