// Copyright © ivanlobanov. All rights reserved.
package middleware

import (
	"context"
	"net/http"
	"time"

	f "github.com/cantylv/online-music-lib/internal/helpers/function"
	mc "github.com/cantylv/online-music-lib/internal/helpers/my/constant"
	"github.com/cantylv/online-music-lib/internal/helpers/recorder"
	"github.com/satori/uuid"
	"go.uber.org/zap"
)

type accessLogStart struct {
	UserAgent      string
	RealIp         string
	ContentLength  int64
	URI            string
	Method         string
	StartTimeHuman string
	RequestId      string
	Logger         *zap.Logger
}

type accessLogEnd struct {
	LatencyMs      int64
	ResponseSize   string // in bytes
	ResponseStatus int
	EndTimeHuman   string
	RequestId      string
	Logger         *zap.Logger
}

// access middleware, который регистрирует начало и конец обработки запроса.
func access(h http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.NewV4().String()
		ctx := context.WithValue(r.Context(), mc.ContextKey(mc.RequestID), requestId)
		r = r.WithContext(ctx)

		rec := recorder.NewResponseWriter(w)
		timeNow := time.Now()
		startLog := accessLogStart{
			UserAgent:      r.UserAgent(),
			RealIp:         r.RemoteAddr,
			ContentLength:  r.ContentLength,
			URI:            r.RequestURI,
			Method:         r.Method,
			StartTimeHuman: f.FormatTime(timeNow),
			RequestId:      requestId,
			Logger:         logger,
		}
		logInitRequest(startLog)
		h.ServeHTTP(rec, r)

		timeEnd := time.Now()
		endLog := accessLogEnd{
			LatencyMs:      timeEnd.Sub(timeNow).Milliseconds(),
			ResponseSize:   w.Header().Get("Content-Length"),
			ResponseStatus: rec.StatusCode,
			EndTimeHuman:   f.FormatTime(timeEnd),
			RequestId:      requestId,
			Logger:         logger,
		}
		logEndRequest(endLog)

	})
}

// logInitRequest регистрирует user-agent, IP-адрес клиента и т. д.
func logInitRequest(data accessLogStart) {
	data.Logger.Info("init request",
		zap.String("user-agent", data.UserAgent),
		zap.String("real-ip", data.RealIp),
		zap.Int64("content-length", data.ContentLength),
		zap.String("uri", data.URI),
		zap.String("method", data.Method),
		zap.String("start-time-human", data.StartTimeHuman),
		zap.String(mc.RequestID, data.RequestId),
	)
}

// LogEndRequest регистрирует задержку в мс, размер ответа и т. д.
func logEndRequest(data accessLogEnd) {
	data.Logger.Info("end of request",
		zap.Int64("latensy-ms", data.LatencyMs),
		zap.String("response-size", data.ResponseSize),
		zap.Int("response-status", data.ResponseStatus),
		zap.String("end-time-human", data.EndTimeHuman),
		zap.String(mc.RequestID, data.RequestId),
	)
}
