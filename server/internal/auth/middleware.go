package auth

import (
	"context"
	"net/http"

	"github.com/neurocome/ine_go/internal/pkg/jwt"
	"github.com/neurocome/ine_go/internal/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			if tokenStr == "" {
				c, err := r.Cookie("belivine")

				// Allow unauthenticated users in
				if err != nil || c == nil {
					next.ServeHTTP(w, r)
					return
				}
				tokenStr = c.Value
			}
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := users.User{Username: username}
			id, err := users.GetUserByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
