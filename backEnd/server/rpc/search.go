package rpc

import (
	"../wordfilter"
	"../constants"
	"../entity"
	"encoding/json"
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
	"net/http"
	"reflect"
	"strconv"
)

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}
	fmt.Printf("Search received: %f %f %s\n", lat, lon, ran)

	// connect to ES
	client, err := elastic.NewClient(elastic.SetURL(ES_URL),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	q := elastic.NewGeoDistanceQuery("location")
	q = q.Distance(ran).Lat(lat).Lon(lon)
	searchResult, err := client.Search().
		Index(INDEX).
		Query(q).
		Pretty(true).
		Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d miiliseconds\n", searchResult.TookInMillis)
	fmt.Printf("Found a total of %d posts\n", searchResult.TotalHits())

	var typ Post
	var ps []Post
	for _, item := range searchResult.Each(reflect.TypeOf(typ)) {
		p := item.(Post) // p = (Post) item
		fmt.Printf("Post by %s: %s at lat %v and lon %v\n",
			p.User, p.Message, p.Location.Lat, p.Location.Lon)

		// Filter restricted words
		if !containsFilteredWords(&p.Message) {
			ps = append(ps, p)
		}
	}

	js, err := json.Marshal(ps)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)

}
