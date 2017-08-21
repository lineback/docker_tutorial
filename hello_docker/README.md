`docker build -t hello_docker:go -f go.dockerfile .`
`docker run --rm -it -p 8080:8080 hello_docker:go`
`curl http://localhost:8080/greeter?name=Steve`