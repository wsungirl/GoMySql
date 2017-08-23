package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsungirl/GoMySql/db"
)

var (
	contextKeyUser  = "user"
	contextKeyToken = "token"
)

var dbGlobal *db.DB

// Setup prepares handlers and sets dbGlobal variable
func Setup(mydb *db.DB) *mux.Router {
	dbGlobal = mydb

	router := setupRouter()

	return router
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	uRtr := r.PathPrefix("/").Subrouter()

	uRtr.HandleFunc("/users", usersHandler).
		Methods("POST")

	uRtr.HandleFunc("/users/{user_id}/auth", usersAuthHandler).
		Methods("GET")

	uRtr.HandleFunc("/users/{user_id}/revoke", withAuth(usersRevokeHandler)).
		Methods("GET")

	uRtr.HandleFunc("/dbs",withAuth(databasesCreate)).
		Methods("POST");

	uRtr.HandleFunc("/dbs",withAuth(databasesList)).
		Methods("GET");

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	return r
}

// notFoundHandler is a placeholder for any request that doesn't have any handler
func notFoundHandler(writer http.ResponseWriter, r *http.Request) {
	returnResult(writer, "Page not found")
}
