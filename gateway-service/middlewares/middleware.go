package middlewares

import (
	"context"
	"gateway/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type contextUserId string
type contextUserRole string

const ContextUserIdKey contextUserId = "userId"
const ContextUserRoleKey contextUserRole = "userRole"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestTarget := r.Method + " " + r.URL.Path

		endpointRoles, err := utils.IsProtectedRoute(requestTarget)

		if err != nil {
			http.Error(w, "Could not find roles for this endpoint", http.StatusInternalServerError)
			return
		}

		if endpointRoles == nil {
			http.Error(w, "Could not find roles for this endpoint", http.StatusInternalServerError)
			return
		}

		if endpointRoles[0] == "" {
			next.ServeHTTP(w, r)
			return
		}

		authToken, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		verifiedToken, err := utils.VerifyJwt(authToken.Value)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := verifiedToken.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Could not convert jwt claims", http.StatusInternalServerError)
			return
		}

		userId, ok := claims["user_id"].(string)

		if !ok {
			http.Error(w, "Could not convert user id from jwt claims to string", http.StatusInternalServerError)
			return
		}

		userRole, ok := claims["user_role"].(string)

		if !ok {
			http.Error(w, "Could not convert user role from jwt claims to string", http.StatusInternalServerError)
			return
		}

		for _, v := range endpointRoles {
			if userRole == v {
				ctx := context.WithValue(r.Context(), ContextUserIdKey, userId)
				ctx = context.WithValue(ctx, ContextUserRoleKey, userRole)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "You do not have the right role to do this", http.StatusUnauthorized)
	})
}
