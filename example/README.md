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
[
  {
    "timestamp": "2024-05-15T21:47:52.353673+09:00",
    "severity": "DEBUG",
    "sourceLocation": {
      "function": "main.handler",
      "file": "/Users/shoyo/10.oss/soglog/example/main.go",
      "line": 26
    },
    "message": "slog debug test",
    "labels": {
      "tenant_id": "sampleTenantID"
    }
  },
  {
    "timestamp": "2024-05-15T21:47:52.354056+09:00",
    "severity": "INFO",
    "sourceLocation": {
      "function": "main.handler",
      "file": "/Users/shoyo/10.oss/soglog/example/main.go",
      "line": 27
    },
    "message": "slog init test",
    "labels": {
      "tenant_id": "sampleTenantID"
    }
  },
  {
    "timestamp": "2024-05-15T21:47:52.354063+09:00",
    "severity": "Warning",
    "sourceLocation": {
      "function": "main.handler",
      "file": "/Users/shoyo/10.oss/soglog/example/main.go",
      "line": 28
    },
    "message": "slog warn test",
    "labels": {
      "tenant_id": "sampleTenantID"
    }
  },
  {
    "timestamp": "2024-05-15T21:47:52.354067+09:00",
    "severity": "ERROR",
    "sourceLocation": {
      "function": "main.handler",
      "file": "/Users/shoyo/10.oss/soglog/example/main.go",
      "line": 29
    },
    "message": "slog error test",
    "labels": {
      "tenant_id": "sampleTenantID"
    },
    "stack_trace": "main.handler(...)\n\t/Users/shoyo/10.oss/soglog/example/main.go:29\nmain.main.func1(...)\n\t/Users/shoyo/10.oss/soglog/example/main.go:48\nnet/http.HandlerFunc.ServeHTTP(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:2166\nnet/http.(*ServeMux).ServeHTTP(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:2683\nnet/http.serverHandler.ServeHTTP(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:3137\nnet/http.(*conn).serve(...)\n\t/Users/shoyo/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.3.darwin-arm64/src/net/http/server.go:2039\n"
  }
]

```
