package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsungirl/GoMySql/model"
)

func usersHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		returnResult(w, "Unsupported method: "+req.Method)
		return
	}

	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		returnResult(w, "Cant't read input")
		return
	}
	defer req.Body.Close()

	var user model.User

	if err = json.Unmarshal(payload, &user); err != nil {
		returnResult(w, "Malformed input: "+string(payload))
		return
	}

	returnResult(w, "")
}

func usersAuthHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["user_id"]

	returnResult(w, "auth")
}

func usersRevokeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["user_id"]

	returnResult(w, "revoke")
}
