package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/s4s7/soglog"
)

const ctxKeyTenantID = "TenantID"
const labelKeyTenantID = "tenant_id"

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
		slog.DebugContext(ctx, "slog debug test")
		slog.InfoContext(ctx, "slog init test")
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
	// initialize slog
	slog.SetDefault(slog.New(soglog.NewCloudLoggingHandler("YourProjectID", true, addUserIDToLabelFiled)))

	mux := http.NewServeMux()
	mux.HandleFunc("/soglog", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		handler(w, r.WithContext(setTenantIDToCtx(ctx, "sampleTenantID")))
	})

	http.ListenAndServe(":8080", mux)
}
