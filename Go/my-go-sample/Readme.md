# my-go-sample

This is a standalone Go sample project. Designed to be easily compiled and executed in both local and containerized environments.

---

## 1. Execution (Run)

### Local Environment
Run the code directly or build a binary in your local Go environment.
```bash
# Option 1: Run directly
$ go run main.go

# Option 2: Build and run
$ go build main.go
$ ./main
```

### Docker Environment

Run the code using a temporary Go container (one-liner).

```bash
$ docker run --rm -v $(pwd):/app -w /app golang:1.22-alpine go run main.go
```

## 2. Testing

### Local Environment

Run all tests in the current directory with verbose output.

```bash
$ go test -v
```

### Docker Environment

Run tests using a temporary Go container.

```bash
$ docker run --rm -v $(pwd):/app -w /app golang:1.22-alpine go test -v
```


## 3. Docker Image Management

Use these commands to build and manage a persistent Docker image for this sample.

### Build

Build the image with a tag for the local registry.

```bash
$ docker build -t localhost:32000/my-go-sample:latest .
```

### Push

Push the image to your local registry.

```bash
$ docker push localhost:32000/my-go-sample:latest
```

### Run from Image

Run the application using the built image.

```bash
$ docker run --rm localhost:32000/my-go-sample:latest
```

