package handler

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/wsungirl/GoMySql/model"
)

func usersHandler(w http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		returnResult(w, "Cant't read input")
		return
	}
	defer req.Body.Close()

	var user model.User

	if err = json.Unmarshal(payload, &user); err != nil {
		returnResult(w, "Malformed input")
		return
	}

	dbUser, err := dbGlobal.GetUserByName(user.Name)
	if err != nil {
		returnResult(w, "Can't retrieve user info")
		return
	}

	if dbUser != nil {
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
		if err != nil {
			returnResult(w, "Wrong credentials")
			return
		}

		dbUser.PasswordHash = ""
		if payload, err = json.Marshal(&dbUser); err != nil {
			returnResult(w, "Can't serialize user: "+err.Error())
			return
		}

		w.Write(payload)
		return
	}

	if err = dbGlobal.AddUser(&user); err != nil {
		returnResult(w, "Can't add user: "+err.Error())
		return
	}

	dbUser, err = dbGlobal.GetUserByName(user.Name)
	if err != nil {
		returnResult(w, "Can't retrieve user info")
		return
	}

	dbUser.PasswordHash = ""

	if payload, err = json.Marshal(&dbUser); err != nil {
		returnResult(w, "Can't send user data")
		return
	}

	w.Write(payload)
}

func usersAuthHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uidStr := vars["user_id"]

	// Parse authorization header and get token
	authHdr := req.Header.Get("Authorization")
	if authHdr == "" {
		returnResult(w, "Empty Authorization header")
		return
	}

	auth := strings.Fields(authHdr)
	if len(auth) != 2 || strings.ToUpper(auth[0]) != "BASIC" {
		returnResult(w, "Wrong header")
		return
	}

	credentials, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		returnResult(w, "Malformed credentials string")
		return
	}

	// Split credentials and get user's name and password
	auth = strings.Split(string(credentials), ":")
	if len(auth) != 2 {
		returnResult(w, "Malformed credentials string")
		return
	}

	userID, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		returnResult(w, "Error parsing user id: "+uidStr)
		return
	}

	user, err := dbGlobal.GetUser(userID)
	if err != nil {
		returnResult(w, "Can't retrieve user with such id: "+err.Error())
		return
	}

	// Compare user's name and password to hash from db
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(auth[1]))
	if user.Name != auth[0] || err != nil {
		returnResult(w, "Wrong credentials")
		return
	}

	// Generate and save access_token
	token := uuid.NewV4()

	session := model.Session{userID, string(token.Bytes())}

	payload, err := json.Marshal(&session)
	if err != nil {
		returnResult(w, "Can't marshal session object")
		return
	}

	w.Write(payload)
}

func usersRevokeHandler(w http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(req)
	// userID := vars["user_id"]

	returnResult(w, "revoke")
}
