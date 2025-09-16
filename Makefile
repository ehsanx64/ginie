run:
	go run .

dev:
	air

dev-deps:
	go install github.com/air-verse/air@latest

deps:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql

build:
	go build -o tmp/ginie .

clean:
	rm -fr tmp


