package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wsungirl/GoMySql/model"
)

func databasesCreate(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(&contextKeyUser).(*model.User)

	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		returnResult(w, "Cant't read input")
		return
	}
	defer req.Body.Close()

	var database model.Database

	if err = json.Unmarshal(payload, &database); err != nil {
		returnResult(w, "Malformed input")
		return
	}

	database.User = *user
	err = dbGlobal.CreateDB(&database)
	if err != nil {
		returnResult(w, "cant create database : "+err.Error())
		return
	}

	returnResult(w, "")
}

func databasesList(w http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(&contextKeyUser).(*model.User)
	dbList, err := dbGlobal.GetUserDatabases(user)
	if err != nil {
		returnResult(w, "cant get dblist: "+err.Error())
		return
	}

	payload, err := json.Marshal(dbList)
	if err != nil {
		returnResult(w, "cant Marshal : "+err.Error())
		return
	}

	w.Write(payload)
}
