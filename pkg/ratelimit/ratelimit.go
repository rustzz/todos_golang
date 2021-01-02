package ratelimit

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	errs "github.com/rustzz/todos/cmd/errors"
)

var apiErrors = errs.GetErrorsData()

// Check : ...
func Check(next http.Handler) http.Handler {
	/*
	 * NewLimiter(кол-во запросов, в какой промежуток времени)
	 */
	limit := tollbooth.NewLimiter(60, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})
	limit.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	limit.SetMethods([]string{"GET", "POST"})
	middle := func(writer http.ResponseWriter, request *http.Request) {
		httpError := tollbooth.LimitByRequest(limit, writer, request)
		if httpError != nil {
			limit.ExecOnLimitReached(writer, request)
			writer.Header().Add("Content-Type", limit.GetMessageContentType())
			writer.WriteHeader(httpError.StatusCode)

			json.NewEncoder(writer).Encode(apiErrors.RateLimitError)
			return
		}
		next.ServeHTTP(writer, request)
	}
	return http.HandlerFunc(middle)
}
