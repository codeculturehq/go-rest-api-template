# Go REST API Template

Scalable webserver written in Go that uses JWT's for authentication and MongoDB to persist data. All endpoints are logically structured through subrouters. The endpoint for sign up is `/auth/signup` and the endpoint to log in is `/auth/login`. Since certain features like enforcing a certain password strength are not implemented yet, please use it with caution.

## Usage

Add an `.env` file according to the `.example.env` file. Install dependencies with the command `go get` and start the server with `go run main.go` in the root folder.