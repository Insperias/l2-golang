package middleware

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

//RecoverMiddleware Logging and recovering after panic
func RecoverMiddleware(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			defer func() {
				if rcv := recover(); rcv != nil {
					err, ok := rcv.(error)
					if !ok {
						err = fmt.Errorf("%v", rcv)
					}

					log.Errorf("[panic recover] %v\n", err)
					resp.WriteHeader(http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(resp, req)
		})
	}
}
