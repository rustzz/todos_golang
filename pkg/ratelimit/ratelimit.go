package ratelimit

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

// Check : ...
func Check(next http.Handler) http.Handler {
	/*
	 * NewLimiter(кол-во запросов, в какой промежуток времени)
	 */
	limit := tollbooth.NewLimiter(3, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})
	limit.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	limit.SetMethods([]string{"GET", "POST"})
	middle := func(writer http.ResponseWriter, request *http.Request) {
		httpError := tollbooth.LimitByRequest(limit, writer, request)
		if httpError != nil {
			limit.ExecOnLimitReached(writer, request)
			writer.Header().Add("Content-Type", limit.GetMessageContentType())
			writer.WriteHeader(httpError.StatusCode)
			writer.Write([]byte(httpError.Message))
			return
		}
		next.ServeHTTP(writer, request)
	}
	return http.HandlerFunc(middle)
}
