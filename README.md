# partylist-server
Shared playlist service

# Devving

## Build Container
```
$ docker build -t byxorna/partylist-server .
```

## Running with containers

```
$ docker run --name partylist-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
$ docker run --name partylist --link partylist-postgres:postgres -d byxorna/partylist-server

```
