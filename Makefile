.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler/send-new-order-to-sns handler/send-new-order-to-sns/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler/products handler/products/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler/products-by-article handler/products-by-article/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler/featured-products handler/featured-products/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
