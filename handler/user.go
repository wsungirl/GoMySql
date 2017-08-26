package handler

import (
	"encoding/base64"
	"encoding/hex"
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

// usersHandler handles POST requests to /users
// if user with provided credentials wasn't found we create one
// if name and password doesn't match we return error
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
		returnResult(w, "Can't retrieve user info: "+err.Error())
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

// usersAuthHandler creates session and returns access_token to use with methods wrapped by withAuth
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

	userID, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		returnResult(w, "Error parsing user id: "+uidStr)
		return
	}

	user, err := dbGlobal.GetUser(uint(userID))
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
	token := hex.EncodeToString(uuid.NewV4().Bytes())
	session := model.Session{User: *user, AccessToken: token}

	if err = dbGlobal.AddSession(&session); err != nil {
		returnResult(w, "Can't add user session")
		return
	}

	payload, err := json.Marshal(&session)
	if err != nil {
		returnResult(w, "Can't marshal session object")
		return
	}

	w.Write(payload)
}

// usersRevokeHandler deletes session from DB
func usersRevokeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		returnResult(w, "Can't parse ID")
		return
	}

	user := req.Context().Value(&contextKeyUser).(*model.User)

	if user.ID != uint(userID) {
		returnResult(w, "Can't revoke access_token of other user")
		return
	}

	token := req.Context().Value(&contextKeyToken).(*string)
	session := model.Session{User: *user, AccessToken: *token}

	if err = dbGlobal.RevokeSession(&session); err != nil {
		returnResult(w, "Can't revore token: "+err.Error())
		return
	}

	returnResult(w, "")
}
