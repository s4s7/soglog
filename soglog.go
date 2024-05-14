package soglog

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"

	"cloud.google.com/go/logging"
	"go.opentelemetry.io/otel/trace"
)

// See: https://cloud.google.com/logging/docs/agent/logging/configuration#special-fields
const (
	keySource       string = "logging.googleapis.com/sourceLocation"
	keyLabel        string = "logging.googleapis.com/labels"
	keyTrace        string = "logging.googleapis.com/trace"
	keySpan         string = "logging.googleapis.com/spanID"
	keyTraceSampled string = "logging.googleapis.com/trace_sampled"
	keyTime         string = "timestamp"
	keyMessage      string = "message"
	keySeverity     string = "severity"
	keyStack        string = "stack_trace"

	traceFmt = "projects/%s/traces/%s"
)

var _ slog.Handler = (*CloudLoggingHandler)(nil)

type CloudLoggingHandler struct {
	handler            slog.Handler
	projectID          string
	isStackTraced      bool
	labelFieldInjector LabelFieldInjector
}

type LabelFieldInjector func(ctx context.Context) (key, value string, found bool)

func NewCloudLoggingHandler(projectID string, isStackTraced bool, labelFieldInjector LabelFieldInjector) *CloudLoggingHandler {
	return &CloudLoggingHandler{
		handler: slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelDebug,
			ReplaceAttr: replaceAttr,
		}),
		projectID:          projectID,
		isStackTraced:      isStackTraced,
		labelFieldInjector: labelFieldInjector,
	}
}

func (h *CloudLoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *CloudLoggingHandler) Handle(ctx context.Context, rec slog.Record) error {

	// set trace info
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.IsValid() {
		rec.AddAttrs(
			slog.String(keyTrace, fmt.Sprintf(traceFmt, h.projectID, spanCtx.TraceID().String())),
			slog.String(keySpan, spanCtx.SpanID().String()),
			slog.Bool(keyTraceSampled, spanCtx.IsSampled()),
		)
	}

	// add customized label
	if h.labelFieldInjector != nil {
		key, value, found := h.labelFieldInjector(ctx)
		if found {
			rec.AddAttrs(slog.Group(keyLabel, slog.String(key, value)))
		}
	}

	// set stack trace
	if h.isStackTraced && rec.Level.String() == slog.LevelError.String() {
		rec.AddAttrs(
			// skip 3 {this func, slog.(*Logger).log, slog.ErrorContext}
			slog.String(keyStack, string(newStackFrames(3))),
		)
	}

	return h.handler.Handle(ctx, rec)
}

func (h *CloudLoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CloudLoggingHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *CloudLoggingHandler) WithGroup(name string) slog.Handler {
	return &CloudLoggingHandler{handler: h.handler.WithGroup(name)}
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		a.Key = keyTime
	case slog.MessageKey:
		a.Key = keyMessage
	case slog.SourceKey:
		a.Key = keySource
	case slog.LevelKey:
		a.Key = keySeverity
		if a.Value.String() == slog.LevelWarn.String() {
			a.Value = slog.StringValue(logging.Warning.String())
		}
	}
	return a
}

func newStackFrames(skip int) []byte {
	const numFrames = 32
	pcs := [numFrames]uintptr{}

	// skip [runtime.Callers, this function]
	n := runtime.Callers(skip+2, pcs[:])

	buf := bytes.Buffer{}
	frames := runtime.CallersFrames(pcs[:n])
	for {
		f, ok := frames.Next()
		if !ok {
			break
		}
		buf.WriteString(f.Function)
		buf.WriteString("(...)\n\t")
		buf.WriteString(f.File)
		buf.Write([]byte{':'})
		buf.WriteString(strconv.Itoa(f.Line))
		buf.Write([]byte{'\n'})
	}

	return buf.Bytes()
}
