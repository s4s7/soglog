package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/s4s7/soglog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

const ctxKeyTenantID = "TenantID"
const labelKeyTenantID = "tenant_id"

var tracer = otel.Tracer("github.com/s4s7/soglog/example")

func setTenantIDToCtx(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, ctxKeyTenantID, tenantID)
}

func getTenantIDFromCtx(ctx context.Context) (string, bool) {
	tenantID, ok := ctx.Value(ctxKeyTenantID).(string)
	return tenantID, ok
}

func handler(_ http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method == http.MethodGet {
		ctx, span := tracer.Start(ctx, "span.test")
		defer span.End()

		// show log at each log level
		slog.DebugContext(ctx, "slog debug test")
		slog.InfoContext(ctx, "slog info test")
		slog.WarnContext(ctx, "slog warn test")
		slog.ErrorContext(ctx, "slog error test")
	}
}

func addUserIDToLabelFiled(ctx context.Context) (key, value string, found bool) {
	tenantID, found := getTenantIDFromCtx(ctx)
	if found {
		return labelKeyTenantID, tenantID, true
	}
	return "", "", false
}

func main() {
	// init logger
	slog.SetDefault(slog.New(soglog.NewCloudLoggingHandler("YourProjectID", true, addUserIDToLabelFiled)))

	// init tracer
	otel.SetTracerProvider(trace.NewTracerProvider())

	mux := http.NewServeMux()
	mux.HandleFunc("/soglog", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		handler(w, r.WithContext(setTenantIDToCtx(ctx, "sampleTenantID")))
	})

	http.ListenAndServe(":8080", mux)
}
