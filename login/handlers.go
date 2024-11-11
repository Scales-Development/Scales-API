package login

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

	if u.Username == "Insert_Username" && u.Password == "Insert_Password" {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("Invalid Credentials/User not found")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid Credentials/User not found")
	}
} 

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "Welcome to Scales-API, Please direct a route to grab data from.")
}