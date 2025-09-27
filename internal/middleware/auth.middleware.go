package middleware

import (
	"context"
	"github/Chidi-creator/go-medic-server/internal/managers"
	"github/Chidi-creator/go-medic-server/internal/services"
	"github/Chidi-creator/go-medic-server/internal/utils"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

// Auth middleware that validates JWT in router
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Extract token from Authorization Header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			managers.JSONresponse(w, http.StatusUnauthorized, utils.ApiResponse{
				Success: false,
				Error:   "Missing Authorisation header",
			})
			return
		}
		//splitting authorisation into two parts and checking validity of structure
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			managers.JSONresponse(w, http.StatusUnauthorized, utils.ApiResponse{
				Success: false,
				Error:   "Invalid Authorisation format",
			})
			return
		}

		tokenstr := parts[1]

		// validating token
		claims, err := services.ValidateToken(tokenstr)
		if err != nil {
			managers.JSONresponse(w, http.StatusUnauthorized, utils.ApiResponse{
				Success: false,
				Error:   "Invalid Token: " + err.Error(),
			})
		}

		//attach claims to context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}

// retrieving user claims from request context

func GetUserFromContext(ctx context.Context) *services.Claims {
	if user, ok := ctx.Value(UserContextKey).(*services.Claims); ok {
		return user
	}
	return nil
}
