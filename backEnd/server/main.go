package main

import (
	"./constants"
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
	"log"
	"net/http"
	// auth
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// transfer string to byte
var mySigningKey = []byte("thoffysecret")

func main() {
	// Create a client
	client, err := elastic.NewClient(elastic.SetURL(ES_URL), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(INDEX).Do()
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		mapping := `{
			"mappings":{
				"post":{
					"properties":{
						"location":{
							"type":"geo_point"
						}
					}
				}
			}
		}`
		_, err = client.CreateIndex(INDEX).Body(mapping).Do()
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	fmt.Println("started-service")

	r := mux.NewRouter()
	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	/*without auth*/
	//http.HandleFunc("/post", handlerPost)
	//http.HandleFunc("/search", handlerSearch)

	/*with auth : use jwtMiddleware to check the token*/
	r.Handle("/post", jwtMiddleware.Handler(http.HandlerFunc(handlerPost))).Methods("POST", "OPTIONS")
	r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(handlerSearch))).Methods("GET", "OPTIONS")

	// these steps needs password, so no auth checking
	r.Handle("/login", http.HandlerFunc(handlerLogIn)).Methods("POST", "OPTIONS")
	r.Handle("/signup", http.HandlerFunc(handlerSignUp)).Methods("POST", "OPTIONS")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
