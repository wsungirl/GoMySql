package handler

import "net/http"



func databasesCreate(w http.ResponseWriter, req *http.Request){

}

func databasesList(w http.ResponseWriter, req *http.Request){
	dbGlobal.GetDBList()
}