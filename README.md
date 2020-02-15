# xdebug-trace-performance-checker

PHP xdebug auto_trace Perfomance checking tool

## Usage

### Download binary

Releases

### Using Golang

```sh
go build -o xdebug-trace-performance-checker main.go
```

```sh
> xdebug-trace-performance-checker debug.xt 0.1
Line: 2010  Duration: 0.17930000000000001
{2.9641 7374504 14 curl_multi_exec() /var/www/vendor/guzzlehttp/guzzle/src/Handler/CurlMultiHandler.php 108}
{2.7848 7374480 13 GuzzleHttp\Handler\CurlMultiHandler->tick() /var/www/vendor/guzzlehttp/guzzle/src/Handler/CurlMultiHandler.php 125}
{2.6678 7374480 12 GuzzleHttp\Handler\CurlMultiHandler->execute() /var/www/vendor/guzzlehttp/promises/src/Promise.php 246}
{2.6678 7374480 11 GuzzleHttp\Promise\Promise->invokeWaitFn() /var/www/vendor/guzzlehttp/promises/src/Promise.php 223}
{2.6678 7374480 10 GuzzleHttp\Promise\Promise->waitIfPending() /var/www/vendor/guzzlehttp/promises/src/Promise.php 267}
{2.6678 7374480 9 GuzzleHttp\Promise\Promise->invokeWaitList() /var/www/vendor/guzzlehttp/promises/src/Promise.php 225}
{2.6678 7374480 8 GuzzleHttp\Promise\Promise->waitIfPending() /var/www/vendor/guzzlehttp/promises/src/Promise.php 62}
{2.6678 7374480 7 GuzzleHttp\Promise\Promise->wait() /var/www/vendor/guzzlehttp/promises/src/EachPromise.php 101}
{2.6677 7374456 6 GuzzleHttp\Promise\EachPromise->GuzzleHttp\Promise\{closure}() /var/www/vendor/guzzlehttp/promises/src/Promise.php 246}
{2.6677 7374456 5 GuzzleHttp\Promise\Promise->invokeWaitFn() /var/www/vendor/guzzlehttp/promises/src/Promise.php 223}
{2.6677 7374456 4 GuzzleHttp\Promise\Promise->waitIfPending() /var/www/vendor/guzzlehttp/promises/src/Promise.php 62}
...
```

## LICENSE

MIT
