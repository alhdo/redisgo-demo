package cors

import "net/http"

type Handler struct {
}

func (ch *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	rw.Header().Set("Access-Control-Max-Age", "3600")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	rw.WriteHeader(http.StatusNoContent)
}

func AddCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow GET, POST, PUT, DELETE methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		// Allow the "Content-Type" header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Allow credentials (if needed)
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests (OPTION method)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue to the next handler
		handler.ServeHTTP(w, r)
	})
}
