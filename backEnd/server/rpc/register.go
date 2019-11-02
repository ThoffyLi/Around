package rpc

import (
	"../user"
	"encoding/json"
	"fmt"
	"net/http"
)

// signup: if suffessful, a new session is created
func handlerSignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on signup request")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var u User
	if err := decoder.Decode(&u); err != nil {
		panic(err)
	}
	if u.Username != "" && u.Password != "" && usernamePattern(u.Username) {
		if addUser(u) {
			fmt.Println("User added!")
			w.Write([]byte("User added!"))
		} else {
			fmt.Println("Failed to add a new user!")
			http.Error(w, "Failed to add a new user!!", http.StatusInternalServerError)
		}
	} else {
		fmt.Println("Invalid password or username!")
		http.Error(w, "Invalid password or username!", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
