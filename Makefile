gorun: 
	go mod tidy
	go run main.go

gobuild:
	go build -o bin/main main.go

gocompile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

gotest:
	cd src/usecases/
	go test ./...

build:
	docker build --tag michaelahli/linkaja:beta .

create:
	docker container create --name linkaja-container -p 8080:8080 -e SERVER_PORT=8080 -e MONGO_DB=LinkAja -e MONGO_URI="mongodb://localhost:27017/" -e CORS_HEADER=x-api-key,Authorization,Content-Type,Origin,Accept,Access-Control-Allow-Headers,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin -e CORS_METHOD=OPTION,GET,PUT,POST,DELETE -e CORS_ORIGIN='*' michaelahli/linkaja:beta

start:
	docker container start linkaja-container

compose :
	docker-compose up -d

stop:
	docker container stop linkaja-container

docker: build compose
