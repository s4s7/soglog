## 1. set up sample code
```sh
go mod tidy
go run main.go
```

## 2. curl
```sh
curl http://localhost:8080/soglog
```

## 3. check log
```json
{
  "timestamp": "2024-05-17T02:04:48.706621+09:00",
  "severity": "DEBUG",
  "logging.googleapis.com/sourceLocation": {
    "function": "main.handler",
    "file": "/Users/shoyo/10.oss/soglog/example/main.go",
    "line": 34
  },
  "message": "slog debug test",
  "logging.googleapis.com/trace": "projects/YourProjectID/traces/fe54d65a6caeb82491c5a44829f7302a",
  "logging.googleapis.com/spanID": "a21bf473806aee5f",
  "logging.googleapis.com/trace_sampled": true,
  "logging.googleapis.com/labels": {
    "tenant_id": "sampleTenantID"
  }
}

{
  "timestamp": "2024-05-17T02:04:48.706872+09:00",
  "severity": "INFO",
  "logging.googleapis.com/sourceLocation": {
    "function": "main.handler",
    "file": "/Users/shoyo/10.oss/soglog/example/main.go",
    "line": 35
  },
  "message": "slog info test",
  "logging.googleapis.com/trace": "projects/YourProjectID/traces/fe54d65a6caeb82491c5a44829f7302a",
  "logging.googleapis.com/spanID": "a21bf473806aee5f",
  "logging.googleapis.com/trace_sampled": true,
  "logging.googleapis.com/labels": {
    "tenant_id": "sampleTenantID"
  }
}

{
  "timestamp": "2024-05-17T02:04:48.706879+09:00",
  "severity": "Warning",
  "logging.googleapis.com/sourceLocation": {
    "function": "main.handler",
    "file": "/Users/shoyo/10.oss/soglog/example/main.go",
    "line": 36
  },
  "message": "slog warn test",
  "logging.googleapis.com/trace": "projects/YourProjectID/traces/fe54d65a6caeb82491c5a44829f7302a",
  "logging.googleapis.com/spanID": "a21bf473806aee5f",
  "logging.googleapis.com/trace_sampled": true,
  "logging.googleapis.com/labels": {
    "tenant_id": "sampleTenantID"
  }
}

{
  "timestamp": "2024-05-17T02:04:48.706904+09:00",
  "severity": "ERROR",
  "logging.googleapis.com/sourceLocation": {
    "function": "main.handler",
    "file": "/Users/shoyo/10.oss/soglog/example/main.go",
    "line": 37
  },
  "message": "slog error test",
  "logging.googleapis.com/trace": "projects/YourProjectID/traces/fe54d65a6caeb82491c5a44829f7302a",
  "logging.googleapis.com/spanID": "a21bf473806aee5f",
  "logging.googleapis.com/trace_sampled": true,
  "logging.googleapis.com/labels": {
    "tenant_id": "sampleTenantID"
  },
  "stack_trace": "main.handler(...)\n\t/Users/shoyo/10.oss/soglog/example/main.go:37\nmain.main.func1(...)\n\t/Users/shoyo/10.oss/soglog/example/main.go:68\nnet/http.HandlerFunc.ServeHTTP(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:2166\nnet/http.(*ServeMux).ServeHTTP(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:2683\nnet/http.serverHandler.ServeHTTP(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:3137\nnet/http.(*conn).serve(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:2039\n"
}

```
