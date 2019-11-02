package rpc

import (
	"../user"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func handlerLogIn(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received one login request")

	w.Header().Set("Content-type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var u User
	if err := decoder.Decode(&u); err != nil {
		panic(err)
	}

	if checkUser(u.Username, u.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		/* Set token claims */
		claims["username"] = u.Username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		/* Sign the token with our secret */
		tokenString, _ := token.SignedString(mySigningKey)

		/* Finally, write the token to the browser window */
		w.Write([]byte(tokenString))
	} else {
		fmt.Println("Invalid password or username.")
		http.Error(w, "Invalid password or username", http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}
