# partylist-server
Shared playlist service

# Devving

## Build Container
```
$ docker build -t byxorna/partylist-server .
```

## Running with containers

```
FILL ME IN

```

## Testing

```
$ godep go build && ./partylist-server -port 6666 -redis-host 192.168.99.101 -redis-port 6379
$ curl -v -XPOST :6666/api/v1/playlist -d '{"name":"test","owner":"gabe"}'
$ curl -v :6666/api/v1/playlist/hjo5me0keaNVwfqjxcgSp6SsJzYDsn
```

# TODO

* [ ] The way I store and retrieve models SUCKS. I hate it and it needs to change
