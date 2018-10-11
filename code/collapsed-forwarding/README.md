# Collapsed Forwarding Example: Request Batcher

This is an example of collapsed forwarding which batches together requests for the same URL. Every five seconds, a `requestBatcher` flushes its batches of requests by handling one request from each batch and giving the same result to the rest of the requests in that batch.

This code was written for the sole purpose of explaining the concept of collapsed forwarding. It is not intended for production use and has known inefficiencies and flaws.

## Getting it running

```bash
$ go build && ./collapsed-forwarding
[server] running on port :8081
[proxy] running on port :8080
```

Once the proxy and server are running, make a request to `localhost:8080/<URL>`. The response should be the provided <URL>:

```bash
$ curl localhost:8080/test
/test
```

Make several requests to the proxy in the background.

```bash
$ curl localhost:8080/test &
$ curl localhost:8080/test &
$ curl localhost:8080/test &
$ curl localhost:8080/test &
```

The application should log each request, find four batched requests, and then make a single request to the server:

```
[proxy] /test
[proxy] /test
[proxy] /test
[proxy] /test
4 batched requests under key "/test"
[server] /test
```
