package handler

import (
	"context"
	"net/http"
	"strings"
)

// withAuth wraps anothed http.HandlerFunc and searches for provided token in DB
// if token wasn't found returns error
func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHdr := r.Header.Get("Authorization")
		if authHdr == "" {
			returnResult(w, "Provide token to process")
			return
		}

		auth := strings.Fields(authHdr)
		if len(auth) != 2 || strings.ToUpper(auth[0]) != "BEARER" {
			returnResult(w, "Wrong header")
			return
		}

		user, err := dbGlobal.GetSessionUser(auth[1])
		if err != nil {
			returnResult(w, "Can't retrieve user")
			return
		}

		ctx := context.WithValue(r.Context(), &contextKeyUser, user)
		ctx = context.WithValue(ctx, &contextKeyToken, &auth[1])

		next(w, r.WithContext(ctx))
	}
}
