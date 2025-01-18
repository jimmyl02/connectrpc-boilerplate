package cors

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

// WithCORS is a middleware wrapper that applies CORS to the given handler.
func WithCORS(h http.Handler, origins []string) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
