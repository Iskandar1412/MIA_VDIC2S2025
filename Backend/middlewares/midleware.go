package middlewares

import "net/http"

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log.Println("panic:", rec)
				http.Error(w, "Error interno", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
