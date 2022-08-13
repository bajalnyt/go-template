## Personal template for new Go projects

![example workflow](https://github.com/bajalnyt/go-template/actions/workflows/go.yml/badge.svg)


What this contains:
* Basic Go code that initializes a server and a route. No additional dependencies.
* A [Makefile](./Makefile). Run `make help` to see descriptions
* A [Dockerfile](./Dockerfile) to containerize the app
```sh
docker build -t myapp .
```
* Go lint rules [.golangci.yml](./.golangci.yml)
* A github workflow for CI checks.
