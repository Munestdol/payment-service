image:
	docker build -t payment-service-image:v1 .

container:
	docker run --name payment-service -p 8080:80 -p 58080:50080/tcp --env-file .env payment-service-image:v1

run:
	go run cmd/main.go
