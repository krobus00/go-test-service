# Go Test Service

test lg 1222
## Build image

```bash
docker build . -t go-test-service 
```

## Run service

### Server service

```bash
docker run -p 3000:3000 --name go-test-service-server go-test-service /app/go-test-service server
```

```bash
make run server
```

### Worker service (example)

```bash
docker run --name go-test-service-worker go-test-service /app/go-test-service worker
```

```bash
make run worker
```

## Unit test

### Run unit test

```bash
make test
```

### See unit test code coverage

```bash
make cover
```
