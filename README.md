# loggen

Write fake web access logs to stdout (at 1 sec interval with jitter):

```
$ go run main.go
78.21.94.134 - - [12/02/2024:09:49:57] "GET / HTTP/1.1" 200 162 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"
78.21.94.132 - - [12/02/2024:09:49:59] "GET /about HTTP/1.1" 301 162 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
78.21.94.130 - - [12/02/2024:09:50:01] "POST /api/v1/users HTTP/1.1" 200 162 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"
13.250.24.4 - - [12/02/2024:09:50:02] "GET /settings HTTP/1.1" 200 162 "-" "Mozilla/5.0 (iPhone13,2; U; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/15E148 Safari/602.1"
25.1.4.135 - - [12/02/2024:09:50:04] "GET /api/v1/settings HTTP/1.1" 200 162 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
25.1.4.142 - - [12/02/2024:09:50:05] "GET / HTTP/1.1" 403 162 "-" "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)"
123.25.44.8 - - [12/02/2024:09:50:06] "POST /api/v1/settings HTTP/1.1" 403 162 "-" "curl/8.4.0"
123.25.44.14 - - [12/02/2024:09:50:06] "GET / HTTP/1.1" 403 162 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"
13.250.24.6 - - [12/02/2024:09:50:08] "GET / HTTP/1.1" 200 162 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"
123.25.44.12 - - [12/02/2024:09:50:08] "GET /api/v1/settings HTTP/1.1" 200 162 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"
13.250.24.14 - - [12/02/2024:09:50:10] "GET / HTTP/1.1" 200 162 "-" "curl/8.4.0"
123.25.44.8 - - [12/02/2024:09:50:11] "POST /api/v1/settings HTTP/1.1" 200 162 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"
123.25.44.7 - - [12/02/2024:09:50:12] "GET /api/v1/settings HTTP/1.1" 200 162 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
78.21.94.135 - - [12/02/2024:09:50:14] "POST /api/v1/users HTTP/1.1" 200 162 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"
```

## Develop

```sh
git clone github.com/kubetail-org/loggen
cd loggen
go run main.go
```

## Build

```sh
docker build -t loggen:latest .
```
