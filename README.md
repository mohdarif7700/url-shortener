# url-shortener
URL Shortener

Use can run the code by using below cmd:
```
$ go run main.go
```

API Endpoints:

- [POST] `127.0.0.1:8080/shorten-url`
```
{
    "message": {
        "originalURL": "google.com",
        "shortenedURL": "https://shorturl.xyz/e8399105"
    },
    "status": "success"
}
```
- [GET] `127.0.0.1:8080/redirect?shortURL="your_short_url"`
```
You will get redirected to the original URL
```
- [GET] `127.0.0.1:8080/metrics`
```
{
    "metrics": {
        "https://google.com" :   7,
        "https://yahoo.com" :    2,
        "https://facebook.com" : 3
    },
    "status": "success"
}
```
