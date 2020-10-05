## Samples for running go coverage

[Coverage](https://www.ory.sh/golang-go-code-coverage-accurate/)

```sh
go test -covermode=atomic -coverprofile=count.out <package file> <test file>

go tool cover -func=count.out
go tool cover -html=count.out
```
