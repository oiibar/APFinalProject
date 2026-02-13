package middleware

import (
	"net/http"
	"os"
	"strings"
)

// CORS adds basic development-friendly CORS headers.
//
// Configure allowed origins via ALLOWED_ORIGINS (comma-separated).
// If ALLOWED_ORIGINS is empty, it allows all origins ("*") for ease of local dev.
func CORS(next http.Handler) http.Handler {
	allowed := strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS"))
	allowAll := allowed == ""
	allowedSet := map[string]struct{}{}
	if !allowAll {
		for _, o := range strings.Split(allowed, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				allowedSet[o] = struct{}{}
			}
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowAll {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else if origin != "" {
			if _, ok := allowedSet[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
			}
		}

		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
