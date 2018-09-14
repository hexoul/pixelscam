# Pixelscam

[![GoDoc](https://godoc.org/github.com/hexoul/pixelscam?status.svg)](https://godoc.org/github.com/hexoul/pixelscam)

> Pixelscam Back-end to mine pixel on DB

Pixelscam Front-end is [here](https://github.com/astrocket/pixelscam.me)

## Contents
- [Build](#build)
- [Test](#test)
- [Deploy](#deploy)
- [License](#license)

## Build
```shell
dep ensure
make
```

## Test
```shell
go test -v
```

## Deploy
Option A. AWS Lambda
1. Set Lambda on AWS
  - Function package: compressed binary file in $GOPATH/src/{repo}/bin
  - Handler: pixelscam (binary file name, it is optional)
  - Runtime: Go 1.x
2. Set API Gateway for cache server on AWS
3. Add API Gateway as Lambda trigger
4. Add CloudWatch Logs

## License
MIT © [hexoul](https://github.com/hexoul)
