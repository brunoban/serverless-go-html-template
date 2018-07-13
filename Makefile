build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/template template/main.go

deploy:
	docker run --rm $$(docker build -q .)