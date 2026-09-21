package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/AndroDeMohawk/notes-api/pkg/auth"
	"github.com/AndroDeMohawk/notes-api/pkg/response"
)

type contextKey string

const UserIDKey contextKey = "userID"

func Authorize(jwtManager *auth.JWTManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.SendJson(w, "Authorization header is empty", http.StatusUnauthorized)
				return
			}

			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				response.SendJson(w, "token is invalid", http.StatusUnauthorized)
				return
			}

			tokenStr := headerParts[1]

			userID, err := jwtManager.ParseToken(tokenStr)
			if err != nil {
				response.SendJson(w, "error to parse token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func GetUserID(ctx context.Context) (int64, bool) {
	userId, ok := ctx.Value(UserIDKey).(int64)
	return userId, ok
}
