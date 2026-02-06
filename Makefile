run:
	@cd src ; go run *.go

build:
	@cd src ; go mod tidy ; go build -o ../bin/inventory *.go

clean:
	@rm ./bin/inventory

image:
	@docker build -t inventory:latest .

start:
	@docker run -d -p 80:80 --name inventory inventory:latest

stop:
	@docker rm -f inventory
