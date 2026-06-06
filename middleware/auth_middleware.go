package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"user-management/model"
	"user-management/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type contextKey string

const AuthClaimsContextKey contextKey = "authClaims"

type errorResponse struct {
	Status  bool        `json:"status"`
	Content interface{} `json:"content"`
	Message string      `json:"message"`
}

func JWTAuthMiddleware(secret string, userService service.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				unauthorized(w, "missing authorization header")
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				unauthorized(w, "authorization header format must be Bearer {token}")
				return
			}

			tokenString := parts[1]
			claims := &model.AuthClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				unauthorized(w, "invalid or expired token")
				return
			}

			if _, err := userService.GetUserByID(claims.UserID); err != nil {
				unauthorized(w, "invalid token user")
				return
			}

			ctx := context.WithValue(r.Context(), AuthClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAuthClaims(ctx context.Context) (*model.AuthClaims, bool) {
	claims, ok := ctx.Value(AuthClaimsContextKey).(*model.AuthClaims)
	return claims, ok
}

func unauthorized(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(errorResponse{Status: false, Content: message, Message: "failed"})
}
