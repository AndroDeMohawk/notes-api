package middlewares

import (
	"context"
	"net/http"
	"uuid"

	"github.com/AndroDeMohawk/notes-api/pkg/Ctx"
)

func SenderReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()

		ctx := context.WithValue(r.Context(), Ctx.Key, id.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
