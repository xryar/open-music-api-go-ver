package middlewares

import (
	"context"
	"net/http"
	"open-music-go/exception"
	"open-music-go/helper"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			panic(exception.NewUnauthorizedError("missing token"))
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			panic(exception.NewUnauthorizedError("invalid token format"))
		}

		userId, err := helper.ValidateJWT(token)
		if err != nil {
			panic(exception.NewUnauthorizedError("invalid token"))
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		next(w, r.WithContext(ctx), p)
	}
}
