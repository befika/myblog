include .env
export

run:
	go run cmd/main.go
test:
	cd Testing/http && godog 


	
