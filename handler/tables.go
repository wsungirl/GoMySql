package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wsungirl/GoMySql/model"
)

func tableCreate(w http.ResponseWriter, req *http.Request) {
	//user := req.Context().Value(&contextKeyUser).(*model.User)

	vars := mux.Vars(req)
	dbID, err := strconv.ParseUint(vars["db_id"], 10, 64)
	if err != nil {
		returnResult(w, "Can't parse db id: "+err.Error())
		return
	}

	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		returnResult(w, "Cant't read input")
		return
	}
	defer req.Body.Close()

	var table model.DBTable

	if err = json.Unmarshal(payload, &table); err != nil {
		returnResult(w, "Malformed input")
		return
	}

	table.DB, err = dbGlobal.GetDB(uint(dbID))

	if err != nil {
		returnResult(w, "Can't select table: "+err.Error())
		return
	}

	err = dbGlobal.CreateTable(&table)

	if err != nil {
		returnResult(w, "Can't create table: "+err.Error())
		return
	}

	returnResult(w, "")
}

func tablesList(w http.ResponseWriter, req *http.Request) {
	// user := req.Context().Value(&contextKeyUser).(*model.User)

	vars := mux.Vars(req)
	dbID, err := strconv.ParseUint(vars["db_id"], 10, 64)
	if err != nil {
		returnResult(w, "Can't parse db id: "+err.Error())
		return
	}

	dbMod, err := dbGlobal.GetDB(uint(dbID))
	if err != nil {
		returnResult(w, "Can't get db info: "+err.Error())
		return
	}

	tables, err := dbGlobal.GetDatabaseTables(dbMod)
	if err != nil {
		returnResult(w, "Can't get db tables: "+err.Error())
		return
	}

	payload, err := json.Marshal(tables)
	if err != nil {
		returnResult(w, "Can't marshal: "+err.Error())
		return
	}

	w.Write(payload)
}
