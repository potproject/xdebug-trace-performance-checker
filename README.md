# xdebug-trace-performance-checker

PHP xdebug auto_trace Perfomance checking tool

## Usage

Using Golang

```sh
go build -o xdebug-trace-performance-checker main.go
```

```sh
xdebug-trace-performance-checker debug.xt 0.1
```

Using Docker

```
docker run -it $(docker build -q .) -v $PWD:/app debug.xt 0.1
```

## LICENSE

MIT
