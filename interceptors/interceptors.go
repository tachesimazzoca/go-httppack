package interceptors

import (
	"net/http"
)

type Interceptor interface {
	Intercept(http.Handler) http.Handler
}
