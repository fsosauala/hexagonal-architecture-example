# Golang Hexagonal Architecture Example

## To use the lambda example
* Compile code with 
  * `GOARCH=amd64 GOOS=linux go build -o ./build/bin/ ./cmd/lambda/main.go`
* Zip it with
  * `zip function.zip build/bin/main`
* Create an AWS lambda with the following configuration
  * Go environment
  * x84_64 architecture
* Update the lambda with the following parameters
  * Controller: build/bin/main
* Upload the zip file to the lambda
* Create an API Gateway
  * API REST type
* Create a POST method
* Hook up with the lambda and activate the integration
* Implement the API
* Play with it :D

## To use the microservice example
* Run it with
  * `go run main application.go -port 8080`
  * Or with any other port you want.
